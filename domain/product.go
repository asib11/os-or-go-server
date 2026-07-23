package domain

// model or entity -> existence ->
type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}
