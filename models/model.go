package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string         `gorm:"type:varchar(100);not null"`
	Observations sql.NullString `gorm:"type:varchar(100);"`
	Price        int            `gorm:"type:int;not null"`
}

type Client struct {
	gorm.Model
	Name       string         `gorm:"type:varchar(100);not null"`
	Comment    sql.NullString `gorm:"type:varchar(100);"`
	Country    string         `gorm:"type:varchar(100);not null"`
	PostalCode string         `gorm:"type:varchar(100);not null"`
}

type Sale struct {
	gorm.Model
	ClientID     Client  `gorm:"type:int;not null;references:ID"`
	ProductID    Product `gorm:"type:int;not null;references:ID"`
	Quantity     uint    `gorm:"type:int;not null"`
	ProductPrice Product `gorm:"references:Price"`
	Income       uint    `gorm:"type:int;not null"`
}
