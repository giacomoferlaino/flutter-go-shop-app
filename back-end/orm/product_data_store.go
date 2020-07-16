package orm

import (
	"github.com/jinzhu/gorm"
)

// ProductDataStore is the product data store
type ProductDataStore struct {
	DB *gorm.DB
}

// GetAll returns all the saved products
func (store *ProductDataStore) GetAll() (*Response, error) {
	products := []Product{}
	connection := store.DB.Find(&products)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: products,
	}
	return response, nil
}

// GetByID returns a product based on its ID
func (store *ProductDataStore) GetByID(id uint) (*Response, error) {
	product := Product{}
	connection := store.DB.First(&product, id)
	response := &Response{Meta: MetaData{Rows: connection.RowsAffected}}
	if connection.RecordNotFound() {
		return response, nil
	}
	if connection.Error != nil {
		return nil, connection.Error
	}
	response.Data = []Product{product}
	return response, nil
}

// Add creates a new produt
func (store *ProductDataStore) Add(item interface{}) (*Response, error) {
	product := item.(Product)
	connection := store.DB.Create(&product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []Product{product},
	}
	return response, nil
}

// DeleteByID removes a product based in its ID
func (store *ProductDataStore) DeleteByID(id uint) (*Response, error) {
	product := Product{Model: Model{ID: id}}
	connection := store.DB.Delete(&product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{Meta: MetaData{Rows: connection.RowsAffected}}
	return response, nil
}

// UpdateByID updates a product based on its ID
func (store *ProductDataStore) UpdateByID(id uint, item interface{}) (*Response, error) {
	product := item.(Product)
	targetProduct := Product{}
	targetProduct.ID = id
	connection := store.DB.Model(&targetProduct).Updates(product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []Product{targetProduct},
	}
	return response, nil
}
