package main

import (
	"flutter_shop_app/api"
	"flutter_shop_app/app"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = initDB()
}

func main() {
	defer db.Close()
	app := app.State{
		Database: db,
	}
	router := http.NewServeMux()

	router.Handle("/auth/", api.NewAuthHandler(app))
	router.Handle("/product", api.NewProductHandler(app))
	router.Handle("/order", api.NewOrderHandler(app))

	log.Fatal(http.ListenAndServe(":8080", router))
}
