package product

import (
	"github.com/jinzhu/gorm"
)

type dataStore struct {
	db *gorm.DB
}

func (store *dataStore) getAll() []Product {
	products := &[]Product{}
	store.db.Find(products)
	return *products
}

func (store *dataStore) getByID(id uint) []Product {
	product := &Product{}
	connection := store.db.First(product, id)
	if connection.RecordNotFound() {
		return []Product{}
	}
	return []Product{*product}
}

func (store *dataStore) add(product Product) *Product {
	store.db.Create(&product)
	return &product
}

func (store *dataStore) deleteByID(id uint) int64 {
	product := Product{}
	product.ID = id
	connection := store.db.Delete(&product)
	return connection.RowsAffected
}

func (store *dataStore) updateByID(id uint, product Product) []Product {
	targetProduct := Product{}
	targetProduct.ID = id
	connection := store.db.Model(&targetProduct).Updates(product)
	if connection.RowsAffected == 0 {
		return []Product{}
	}
	return []Product{targetProduct}
}
