package orm

// Order data model
type Order struct {
	Model
	Amount   string    `json:"amount"`
	Products []Product `json:"products"`
	DateTime string    `json:"dateTime"`
}
