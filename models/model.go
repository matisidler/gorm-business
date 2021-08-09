package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string         `gorm:"type:varchar(100);not null"`
	Observations sql.NullString `gorm:"type:varchar(100);"`
	Price        int            `gorm:"type:int"`
	Sale         Sale
}

type Client struct {
	gorm.Model
	Name       string         `gorm:"type:varchar(100);not null"`
	Comment    sql.NullString `gorm:"type:varchar(100);"`
	Country    string         `gorm:"type:varchar(100);not null"`
	PostalCode string         `gorm:"type:varchar(100);not null"`
	Sale       Sale
}

type Sale struct {
	gorm.Model
	ClientID     uint `gorm:"type:int;not null;"`
	ProductID    uint `gorm:"type:int;not null;"`
	ProductPrice uint `gorm:"type:int;not null;"`
	Quantity     uint `gorm:"type:int;not null"`
	Income       uint `gorm:"not null;default:(product_price*quantity)"`
}
