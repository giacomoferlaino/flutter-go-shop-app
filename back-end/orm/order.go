package orm

// Order data model
type Order struct {
	Model
	Amount   string    `json:"amount"`
	Products []Product `json:"products" gorm:"many2many:order_products;"`
	DateTime string    `json:"dateTime"`
}
