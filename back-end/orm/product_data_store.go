package orm

import (
	"github.com/jinzhu/gorm"
)

// ProductDataStore is the product data store
type ProductDataStore struct {
	DB *gorm.DB
}

// GetAll returns all the saved products
func (store *ProductDataStore) GetAll() (interface{}, int64, error) {
	products := []Product{}
	connection := store.DB.Find(&products)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return products, connection.RowsAffected, nil
}

// GetByID returns a product based on its ID
func (store *ProductDataStore) GetByID(id uint) (interface{}, int64, error) {
	product := Product{}
	connection := store.DB.First(&product, id)
	if connection.RecordNotFound() {
		return nil, connection.RowsAffected, nil
	}
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []Product{product}, connection.RowsAffected, nil
}

// Add creates a new product
func (store *ProductDataStore) Add(item interface{}) (interface{}, int64, error) {
	product := item.(*Product)
	connection := store.DB.Create(product)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []Product{*product}, connection.RowsAffected, nil
}

// DeleteByID removes a product based in its ID
func (store *ProductDataStore) DeleteByID(id uint) (int64, error) {
	product := Product{Model: Model{ID: id}}
	connection := store.DB.Delete(&product)
	if connection.Error != nil {
		return connection.RowsAffected, connection.Error
	}
	return connection.RowsAffected, nil
}

// UpdateByID updates a product based on its ID
func (store *ProductDataStore) UpdateByID(id uint, item interface{}) (interface{}, int64, error) {
	product := item.(*Product)
	targetProduct := Product{}
	targetProduct.ID = id
	connection := store.DB.Model(&targetProduct).Updates(*product)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []Product{targetProduct}, connection.RowsAffected, nil
}
