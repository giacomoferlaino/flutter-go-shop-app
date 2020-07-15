package product

import (
	"github.com/jinzhu/gorm"
)

type dataStore struct {
	db *gorm.DB
}

func (store *dataStore) getAll() ([]Product, error) {
	products := &[]Product{}
	connection := store.db.Find(products)
	if connection.Error != nil {
		return nil, connection.Error
	}
	return *products, nil
}

func (store *dataStore) getByID(id uint) ([]Product, error) {
	product := &Product{}
	connection := store.db.First(product, id)
	if connection.RecordNotFound() {
		return []Product{}, nil
	}
	if connection.Error != nil {
		return nil, connection.Error
	}
	return []Product{*product}, nil
}

func (store *dataStore) add(product Product) (*Product, error) {
	connection := store.db.Create(&product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	return &product, nil
}

func (store *dataStore) deleteByID(id uint) (int64, error) {
	product := Product{}
	product.ID = id
	connection := store.db.Delete(&product)
	if connection.Error != nil {
		return 0, connection.Error
	}
	return connection.RowsAffected, nil
}

func (store *dataStore) updateByID(id uint, product Product) ([]Product, error) {
	targetProduct := Product{}
	targetProduct.ID = id
	connection := store.db.Model(&targetProduct).Updates(product)
	if connection.Error != nil {
		return nil, connection.Error
	}
	if connection.RowsAffected == 0 {
		return []Product{}, nil
	}
	return []Product{targetProduct}, nil
}
