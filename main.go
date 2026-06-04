package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // Create a new ServeMux to handle routes. its called router in other languages

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /products", http.HandlerFunc(createProducts))

	globalRouter := globalRouter(mux)

	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	ProductList = []Product{
		{ID: 1, Title: "Product 1", Description: "Description of Product 1", Price: 19.99, ImageUrl: "https://example.com/product1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description of Product 2", Price: 29.99, ImageUrl: "https://example.com/product2.jpg"},
		{ID: 3, Title: "Product 3", Description: "Description of Product 3", Price: 39.99, ImageUrl: "https://example.com/product3.jpg"},
	}
}