package orm

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jinzhu/gorm"
)

// OrderDataStore is the order data store
type OrderDataStore struct {
	DB *gorm.DB
}

// ParseJSON parses a JSON into a value
func (store *OrderDataStore) ParseJSON(reqBody io.ReadCloser) (interface{}, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return nil, err
	}
	order := &Order{}
	err = json.Unmarshal(body, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetAll returns all the saved orders
func (store *OrderDataStore) GetAll() (*Response, error) {
	orders := []Order{}
	connection := store.DB.Preload("Products").Find(&orders)
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
	}
	if connection.RecordNotFound() {
		return response, nil
	}
	if connection.Error != nil {
		return nil, connection.Error
	}
	response.Data = orders
	return response, nil
}

// GetByID returns a order based on its ID
func (store *OrderDataStore) GetByID(id uint) (*Response, error) {
	order := Order{}
	connection := store.DB.Preload("Products").First(&order, id)
	response := &Response{Meta: MetaData{Rows: connection.RowsAffected}}
	if connection.RecordNotFound() {
		return response, nil
	}
	if connection.Error != nil {
		return nil, connection.Error
	}
	response.Data = []Order{order}
	return response, nil
}

// Add creates a new order
func (store *OrderDataStore) Add(item interface{}) (*Response, error) {
	order := item.(*Order)
	connection := store.DB.Create(order)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []Order{*order},
	}
	return response, nil
}

// DeleteByID removes a order based in its ID
func (store *OrderDataStore) DeleteByID(id uint) (*Response, error) {
	order := Order{Model: Model{ID: id}}
	connection := store.DB.Delete(&order)
	if connection.Error != nil {
		return nil, connection.Error
	}
	connection.Model(&order).Association("Products").Clear()
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{Meta: MetaData{Rows: connection.RowsAffected}}
	return response, nil
}

// UpdateByID updates a order based on its ID
func (store *OrderDataStore) UpdateByID(id uint, item interface{}) (*Response, error) {
	order := item.(*Order)
	order.ID = id
	connection := store.DB.Save(order)
	connection.Model(order).Association("Products").Replace(order.Products)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []Order{*order},
	}
	return response, nil
}
