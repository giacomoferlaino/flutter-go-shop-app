package product

import (
	"flutter_shop_app/orm"

	"github.com/jinzhu/gorm"
)

type dataStore struct {
	db *gorm.DB
}

func (store *dataStore) getAll() (*orm.Response, error) {
	products := []Product{}
	connection := store.db.Find(&products)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &orm.Response{
		Meta: orm.MetaData{Rows: connection.RowsAffected},
		Data: products,
	}
	return response, nil
}

func (store *dataStore) getByID(id uint) (*orm.Response, error) {
	product := Product{}
	connection := store.db.First(&product, id)
	response := &orm.Response{Meta: orm.MetaData{Rows: connection.RowsAffected}}
	if connection.RecordNotFound() {
		return response, nil
	}
	if connection.Error != nil {
		return nil, connection.Error
	}
	response.Data = []Product{product}
	return response, nil
}

func (store *dataStore) add(product Product) (*orm.Response, error) {
	connection := store.db.Create(&product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &orm.Response{
		Meta: orm.MetaData{Rows: connection.RowsAffected},
		Data: []Product{product},
	}
	return response, nil
}

func (store *dataStore) deleteByID(id uint) (*orm.Response, error) {
	product := Product{Model: orm.Model{ID: id}}
	connection := store.db.Delete(&product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &orm.Response{Meta: orm.MetaData{Rows: connection.RowsAffected}}
	return response, nil
}

func (store *dataStore) updateByID(id uint, product Product) (*orm.Response, error) {
	targetProduct := Product{}
	targetProduct.ID = id
	connection := store.db.Model(&targetProduct).Updates(product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &orm.Response{
		Meta: orm.MetaData{Rows: connection.RowsAffected},
		Data: []Product{targetProduct},
	}
	return response, nil
}
