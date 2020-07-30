package orm

import (
	"github.com/jinzhu/gorm"
)

// UserDataStore is the user data store
type UserDataStore struct {
	DB *gorm.DB
}

// GetAll returns all the saved users
func (store *UserDataStore) GetAll(preloading bool) (interface{}, int64, error) {
	users := []User{}
	connection := store.DB
	if preloading {
		connection = connection.Preload("Products").Preload("FavoriteProducts")
	}
	connection = connection.Find(&users)
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return users, connection.RowsAffected, nil
}

// GetByID returns a user based on its ID
func (store *UserDataStore) GetByID(id uint, preloading bool) (interface{}, int64, error) {
	user := User{}
	connection := store.DB
	if preloading {
		connection = connection.Preload("Products").Preload("FavoriteProducts")
	}
	connection = connection.First(&user, id)
	if connection.RecordNotFound() {
		return nil, connection.RowsAffected, nil
	}
	if connection.Error != nil {
		return nil, connection.RowsAffected, connection.Error
	}
	return []User{user}, connection.RowsAffected, nil
}

// GetByEmail returns a user based on its ID
func (store *UserDataStore) GetByEmail(email string, preloading bool) (interface{}, int64, error) {
	user := User{}
	connection := store.DB
	if preloading {
		connection = connection.Preload("Products").Preload("FavoriteProducts")
	}
	connection = connection.Where(&User{Email: email}).First(&user)
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

// AddProduct adds a new product to a given user
func (store *UserDataStore) AddProduct(userID uint, productID uint) error {
	data, _, err := store.GetByID(userID, false)
	if err != nil {
		return err
	}
	users, _ := data.([]User)
	connection := store.DB.Model(&users[0]).Association("Products").Append(Product{Model: Model{ID: productID}})
	return connection.Error
}

// RemoveProduct removes a product from a given user
func (store *UserDataStore) RemoveProduct(userID uint, productID uint) error {
	data, _, err := store.GetByID(userID, false)
	if err != nil {
		return err
	}
	users, _ := data.([]User)
	connection := store.DB.Model(&users[0]).Association("Products").Delete(Product{Model: Model{ID: productID}})
	return connection.Error
}

// AddFavoriteProduct adds a new product to a given user
func (store *UserDataStore) AddFavoriteProduct(userID uint, productID uint) error {
	data, _, err := store.GetByID(userID, false)
	if err != nil {
		return err
	}
	users, _ := data.([]User)
	connection := store.DB.Model(&users[0]).Association("FavoriteProducts").Append(Product{Model: Model{ID: productID}})
	return connection.Error
}

// RemoveFavoriteProduct removes a product from a given user
func (store *UserDataStore) RemoveFavoriteProduct(userID uint, productID uint) error {
	data, _, err := store.GetByID(userID, false)
	if err != nil {
		return err
	}
	users, _ := data.([]User)
	connection := store.DB.Model(&users[0]).Association("FavoriteProducts").Delete(Product{Model: Model{ID: productID}})
	return connection.Error
}
