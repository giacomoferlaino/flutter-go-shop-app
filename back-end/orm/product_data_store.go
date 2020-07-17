package orm

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jinzhu/gorm"
)

// ProductDataStore is the product data store
type ProductDataStore struct {
	DB *gorm.DB
}

// ParseJSON parses a JSON into a value
func (store *ProductDataStore) ParseJSON(reqBody io.ReadCloser) (interface{}, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return nil, err
	}
	product := &Product{}
	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}
	return product, nil
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

// Add creates a new product
func (store *ProductDataStore) Add(item interface{}) (*Response, error) {
	product := item.(*Product)
	connection := store.DB.Create(product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []Product{*product},
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
	product := item.(*Product)
	targetProduct := Product{}
	targetProduct.ID = id
	connection := store.DB.Model(&targetProduct).Updates(*product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []Product{targetProduct},
	}
	return response, nil
}
