package main

import (
	"flutter_shop_app/app"
	"flutter_shop_app/orm"
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
	productHandler := NewHandler(app, &orm.ProductDataStore{DB: app.Database})
	orderHandler := NewHandler(app, &orm.OrderDataStore{DB: app.Database})
	router := httprouter.New()
	router.GET("/product", productHandler.Get)
	router.POST("/product", productHandler.Post)
	router.GET("/product/:id", productHandler.GetByID)
	router.PUT("/product/:id", productHandler.UpdateByID)
	router.DELETE("/product/:id", productHandler.DeleteByID)

	router.GET("/order", orderHandler.Get)
	router.POST("/order", orderHandler.Post)
	router.GET("/order/:id", orderHandler.GetByID)
	router.PUT("/order/:id", orderHandler.UpdateByID)
	router.DELETE("/order/:id", orderHandler.DeleteByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
