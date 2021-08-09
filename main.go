package main

import (
	"github.com/matisidler/gorm-business/models"
	"github.com/matisidler/gorm-business/storage"
)

func main() {
	driver := storage.MySQL // driver := storage.MySQL
	storage.NewConnection(driver)

	storage.DB().AutoMigrate(&models.Product{}, &models.Client{}, &models.Sale{})
}
