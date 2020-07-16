package orm

// DataStore contains the methods that all the data store has to implement
type DataStore interface {
	GetAll() (*Response, error)
	GetByID(id uint) (*Response, error)
	Add(item interface{}) (*Response, error)
	DeleteByID(id uint) (*Response, error)
	UpdateByID(id uint, item interface{}) (*Response, error)
}
