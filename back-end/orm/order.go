package orm

// Order data model
type Order struct {
	Model
	Amount    float64    `json:"amount"`
	CartItems []CartItem `json:"cartItems" gorm:"many2many:order_cartItems"`
	DateTime  string     `json:"dateTime"`
}
