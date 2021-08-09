package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

type Driver string

func NewConnection(d Driver) {
	switch d {
	case MySQL:
		newMySqlDB()
	case Postgres:
		newPqDB()
	default:
		log.Fatal("Parametro no valido")
	}

}
func newPqDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/business?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open DB %v", err)
		}

		fmt.Println("Conectado a Postgres.")
	})
	return db
}

func newMySqlDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/business?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open DB %v", err)
		}

		fmt.Println("Conectado a MySQL.")
	})
	return db
}

func StringToNull(s string) sql.NullString {
	var nullString sql.NullString
	if s == "" {
		nullString.Valid = false
	} else {
		nullString.Valid = true
		nullString.String = s
	}
	return nullString
}

func DB() *gorm.DB {
	return db
}
