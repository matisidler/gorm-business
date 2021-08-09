package main

import (
	"github.com/matisidler/gorm-business/models"
	"github.com/matisidler/gorm-business/storage"
)

func main() {
	driver := storage.MySQL // driver := storage.MySQL
	storage.NewConnection(driver)
	/*
		//Migrating Tables
		storage.DB().AutoMigrate(&models.Product{}, &models.Client{}, &models.Sale{})

		//Inserting values
		product1 := models.Product{
			Name:         "Curso Golang",
			Price:        350,
			//Converting strings to nulls and nulls to strings:
			Observations: storage.StringToNull("Curso nuevo"),
		}
		storage.DB().Create(&product1)

		client1 := models.Client{
			Name:       "Juan",
			Country:    "Argentina",
			PostalCode: "5900",
			Comment:    storage.StringToNull("Cliente nuevo"),
		}
		storage.DB().Create(&client1)

		sale1 := models.Sale{
			ClientID:  1,
			ProductID: 2,
			Quantity:  4,
		}
		storage.DB().Create(&sale1) */

	models.ReadSale(1)

}
