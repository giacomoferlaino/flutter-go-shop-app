package orm

import (
	"flutter_shop_app/api"
	"io"

	"github.com/jinzhu/gorm"
)

// UserDataStore is the user data store
type UserDataStore struct {
	DB *gorm.DB
}

// ParseJSON parses a JSON into a value
func (store *UserDataStore) ParseJSON(reqBody io.ReadCloser) (interface{}, error) {
	return api.ParseJSON(reqBody, &User{})
}

// GetAll returns all the saved users
func (store *UserDataStore) GetAll() (interface{}, int64, error) {
	users := []User{}
	connection := store.DB.Find(&users)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return users, connection.RowsAffected, nil
}

// GetByID returns a user based on its ID
func (store *UserDataStore) GetByID(id uint) (interface{}, int64, error) {
	user := User{}
	connection := store.DB.First(&user, id)
	if connection.RecordNotFound() {
		return nil, connection.RowsAffected, nil
	}
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []User{user}, connection.RowsAffected, nil
}

// Add creates a new user
func (store *UserDataStore) Add(item interface{}) (interface{}, int64, error) {
	user := item.(*User)
	connection := store.DB.Create(user)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []User{*user}, connection.RowsAffected, nil
}

// DeleteByID removes a user based in its ID
func (store *UserDataStore) DeleteByID(id uint) (int64, error) {
	user := User{Model: Model{ID: id}}
	connection := store.DB.Delete(&user)
	if connection.Error != nil {
		return connection.RowsAffected, connection.Error
	}
	return connection.RowsAffected, nil
}

// UpdateByID updates a user based on its ID
func (store *UserDataStore) UpdateByID(id uint, item interface{}) (interface{}, int64, error) {
	user := item.(*User)
	targetUser := User{}
	targetUser.ID = id
	connection := store.DB.Model(&targetUser).Updates(*user)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []User{targetUser}, connection.RowsAffected, nil
}
