package product

// Product data model
type Product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
	IsFavorite  bool    `json:"isFavorite"`
}
