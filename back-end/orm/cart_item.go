package orm

// CartItem data model
type CartItem struct {
	Model
	ProductID uint
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
}
