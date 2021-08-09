package models

import (
	"fmt"

	"github.com/matisidler/gorm-business/storage"
	"gorm.io/gorm"
)

//-------PRODUCTS METHODS-------

//Method used to update "product_price" and "income" from table "sales"
func (s *Sale) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&s).Update("product_price", tx.Model(&Product{}).Select("price").Where("sales.product_id = products.id"))
	tx.Model(&s).Update("income", gorm.Expr("product_price * quantity"))
	return
}

//Method used to read a product, introducing an ID
func ReadProduct(id uint) {
	var myProduct = Product{}
	storage.DB().First(&myProduct, id)
	fmt.Printf("ID: %+v | Name: %+v | Price: %+v | Observations: %+v | CreatedAt: %+v | UpdatedAt: %+v | DeletedAt: %+v \n", myProduct.ID, myProduct.Name, myProduct.Price, myProduct.Observations.String, myProduct.CreatedAt, myProduct.UpdatedAt, myProduct.DeletedAt.Time)
}

//Method used to read all products.
func ReadProducts() {
	var myProduct = make([]Product, 0)
	storage.DB().Find(&myProduct)
	for _, product := range myProduct {
		fmt.Printf("ID: %+v | Name: %+v | Price: %+v | Observations: %+v | CreatedAt: %+v | UpdatedAt: %+v | DeletedAt: %+v \n\n", product.ID, product.Name, product.Price, product.Observations.String, product.CreatedAt, product.UpdatedAt, product.DeletedAt.Time)
	}
}

func ReadClients() {
	var myClient = make([]Client, 0)
	storage.DB().Find(&myClient)
	for _, client := range myClient {
		fmt.Printf("ID: %+v | Name: %+v | Country: %+v | PostalCode: %+v | Comment: %+v | CreatedAt: %+v | UpdatedAt: %+v | DeletedAt: %+v \n\n", client.ID, client.Name, client.Country, client.PostalCode, client.Comment.String, client.CreatedAt, client.UpdatedAt, client.DeletedAt.Time)
	}
}

func ReadClient(id uint) {
	var myClient = Client{}
	storage.DB().First(&myClient, id)
	fmt.Printf("ID: %+v | Name: %+v | Country: %+v | PostalCode: %+v | Comment: %+v | CreatedAt: %+v | UpdatedAt: %+v | DeletedAt: %+v \n", myClient.ID, myClient.Name, myClient.Country, myClient.PostalCode, myClient.Comment.String, myClient.CreatedAt, myClient.UpdatedAt, myClient.DeletedAt.Time)
}

func ReadSales() {
	var mySale = make([]Sale, 0)
	storage.DB().Find(&mySale)
	for _, sale := range mySale {
		fmt.Printf("ID: %+v | ClientID: %+v | ProductID: %+v | ProductPrice: %+v | Quantity: %+v | Income: %+v | CreatedAt: %+v | UpdatedAt: %+v | DeletedAt: %+v \n\n", sale.ID, sale.ClientID, sale.ProductID, sale.ProductPrice, sale.Quantity, sale.Income, sale.CreatedAt, sale.UpdatedAt, sale.DeletedAt.Time)
	}
}

func ReadSale(id uint) {
	var sale = Sale{}
	storage.DB().First(&sale, id)
	fmt.Printf("ID: %+v | ClientID: %+v | ProductID: %+v | ProductPrice: %+v | Quantity: %+v | Income: %+v | CreatedAt: %+v | UpdatedAt: %+v | DeletedAt: %+v \n", sale.ID, sale.ClientID, sale.ProductID, sale.ProductPrice, sale.Quantity, sale.Income, sale.CreatedAt, sale.UpdatedAt, sale.DeletedAt.Time)

}
