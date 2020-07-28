package orm

// User data model
type User struct {
	Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Passwod   string `json:"password"`
}
