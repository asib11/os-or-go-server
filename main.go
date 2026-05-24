package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
}

var ProductList []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm Asib and I am learning Go!")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendData(w, ProductList, http.StatusOK)
}

func postProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)

	sendData(w, newProduct, http.StatusCreated)

}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux() // Create a new ServeMux to handle routes
	
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", postProducts)

	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", mux)
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