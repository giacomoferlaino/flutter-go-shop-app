package orm

import "io"

// DataStore contains the methods that all the data store has to implement
type DataStore interface {
	ParseJSON(reqBody io.ReadCloser) (interface{}, error)
	GetAll() (*Response, error)
	GetByID(id uint) (*Response, error)
	Add(item interface{}) (*Response, error)
	DeleteByID(id uint) (*Response, error)
	UpdateByID(id uint, item interface{}) (*Response, error)
}
