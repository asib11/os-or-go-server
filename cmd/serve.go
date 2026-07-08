package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // Create a new ServeMux to handle routes. its called router in other languages

	loggerMiddleware := middleware.Logger(http.HandlerFunc(handlers.TestHandler))
	mux.Handle("GET /test", loggerMiddleware)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProducts))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
