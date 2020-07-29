package orm

import (
	"github.com/jinzhu/gorm"
)

// OrderDataStore is the order data store
type OrderDataStore struct {
	DB *gorm.DB
}

// GetAll returns all the saved orders
func (store *OrderDataStore) GetAll() (interface{}, int64, error) {
	orders := []Order{}
	connection := store.DB.Preload("CartItems.Product").Find(&orders)
	if connection.RecordNotFound() {
		return nil, connection.RowsAffected, nil
	}
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return orders, connection.RowsAffected, nil
}

// GetByID returns a order based on its ID
func (store *OrderDataStore) GetByID(id uint) (interface{}, int64, error) {
	order := Order{}
	connection := store.DB.Preload("CartItems.Product").First(&order, id)
	if connection.RecordNotFound() {
		return nil, connection.RowsAffected, nil
	}
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []Order{order}, connection.RowsAffected, nil
}

// Add creates a new order
func (store *OrderDataStore) Add(item interface{}) (interface{}, int64, error) {
	order := item.(*Order)
	connection := store.DB.Create(order)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []Order{*order}, connection.RowsAffected, nil
}

// DeleteByID removes a order based in its ID
func (store *OrderDataStore) DeleteByID(id uint) (int64, error) {
	order := Order{Model: Model{ID: id}}
	connection := store.DB.Delete(&order)
	if connection.Error != nil {
		return connection.RowsAffected, connection.Error
	}
	connection.Model(&order).Association("CartItems").Clear()
	if connection.Error != nil {
		return connection.RowsAffected, connection.Error
	}
	return connection.RowsAffected, nil
}

// UpdateByID updates a order based on its ID
func (store *OrderDataStore) UpdateByID(id uint, item interface{}) (interface{}, int64, error) {
	order := item.(*Order)
	order.ID = id
	connection := store.DB.Save(order)
	connection.Model(order).Association("CartItems").Replace(order.CartItems)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []Order{*order}, connection.RowsAffected, nil
}
