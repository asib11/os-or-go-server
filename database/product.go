package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
}

func init() {
	ProductList = []Product{
		{ID: 1, Title: "Product 1", Description: "Description of Product 1", Price: 19.99, ImageUrl: "https://example.com/product1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description of Product 2", Price: 29.99, ImageUrl: "https://example.com/product2.jpg"},
		{ID: 3, Title: "Product 3", Description: "Description of Product 3", Price: 39.99, ImageUrl: "https://example.com/product3.jpg"},
	}
}
