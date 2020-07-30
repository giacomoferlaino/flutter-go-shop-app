package orm

// User data model
type User struct {
	Model
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email" gorm:"unique"`
	Password         string    `json:"password"`
	Products         []Product `json:"products" gorm:"many2many:users_products"`
	FavoriteProducts []Product `json:"favoriteProducts" gorm:"many2many:users_favoriteProducts"`
}
