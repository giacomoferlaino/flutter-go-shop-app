package main

import (
	"flutter_shop_app/product"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatalf("Error while connecting to database.\n Error: %v", err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}
