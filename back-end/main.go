package main

import (
	"flutter_shop_app/api"
	"flutter_shop_app/app"
	"flutter_shop_app/auth"
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
	jwtManager := auth.NewJwtManager("secret_key")

	router.Handle("/auth/", api.NewAuthHandler(app, jwtManager))
	router.Handle("/user/", api.NewUserHandler(app, jwtManager))
	router.Handle("/product", api.NewProductHandler(app, jwtManager))
	router.Handle("/order", api.NewOrderHandler(app, jwtManager))

	log.Fatal(http.ListenAndServe(":8080", router))
}
