package orm

// DataStore contains the methods that all the data store has to implement
type DataStore interface {
	GetAll() (interface{}, int64, error)
	GetByID(id uint) (interface{}, int64, error)
	Add(item interface{}) (interface{}, int64, error)
	DeleteByID(id uint) (int64, error)
	UpdateByID(id uint, item interface{}) (interface{}, int64, error)
}
