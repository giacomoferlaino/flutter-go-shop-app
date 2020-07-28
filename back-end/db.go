package main

import (
	"log"

	"flutter_shop_app/orm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatalf("Error while connecting to database.\n Error: %v", err)
	}

	db.AutoMigrate(&orm.Product{})
	db.AutoMigrate(&orm.Order{})
	db.AutoMigrate(&orm.CartItem{})
	db.AutoMigrate(&orm.User{})

	return db
}
