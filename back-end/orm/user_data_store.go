package orm

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jinzhu/gorm"
)

// UserDataStore is the user data store
type UserDataStore struct {
	DB *gorm.DB
}

// ParseJSON parses a JSON into a value
func (store *UserDataStore) ParseJSON(reqBody io.ReadCloser) (interface{}, error) {
	body, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll returns all the saved users
func (store *UserDataStore) GetAll() (*Response, error) {
	users := []User{}
	connection := store.DB.Find(&users)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: users,
	}
	return response, nil
}

// GetByID returns a user based on its ID
func (store *UserDataStore) GetByID(id uint) (*Response, error) {
	user := User{}
	connection := store.DB.First(&user, id)
	response := &Response{Meta: MetaData{Rows: connection.RowsAffected}}
	if connection.RecordNotFound() {
		return response, nil
	}
	if connection.Error != nil {
		return nil, connection.Error
	}
	response.Data = []User{user}
	return response, nil
}

// Add creates a new user
func (store *UserDataStore) Add(item interface{}) (*Response, error) {
	user := item.(*User)
	connection := store.DB.Create(user)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []User{*user},
	}
	return response, nil
}

// DeleteByID removes a user based in its ID
func (store *UserDataStore) DeleteByID(id uint) (*Response, error) {
	user := User{Model: Model{ID: id}}
	connection := store.DB.Delete(&user)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{Meta: MetaData{Rows: connection.RowsAffected}}
	return response, nil
}

// UpdateByID updates a user based on its ID
func (store *UserDataStore) UpdateByID(id uint, item interface{}) (*Response, error) {
	user := item.(*User)
	targetUser := User{}
	targetUser.ID = id
	connection := store.DB.Model(&targetUser).Updates(*user)
	if connection.Error != nil {
		return nil, connection.Error
	}
	response := &Response{
		Meta: MetaData{Rows: connection.RowsAffected},
		Data: []User{targetUser},
	}
	return response, nil
}
