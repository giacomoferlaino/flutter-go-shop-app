package orm

// Product data model
type Product struct {
	Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
	IsFavorite  *bool   `json:"isFavorite"`
}
