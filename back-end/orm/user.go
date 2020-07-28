package orm

// User data model
type User struct {
	Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Passwod   string `json:"password"`
}
