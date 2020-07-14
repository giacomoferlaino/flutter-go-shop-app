package main

import (
	"flutter_shop_app/app"
	"flutter_shop_app/product"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
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
	productHandler := product.NewHandler(app)
	router := httprouter.New()
	router.GET("/product", productHandler.Get)
	router.POST("/product", productHandler.Post)
	router.GET("/product/:id", productHandler.GetByID)
	router.PUT("/product/:id", productHandler.UpdateByID)
	router.DELETE("/product/:id", productHandler.DeleteByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
