package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // Create a new ServeMux to handle routes. its called router in other languages
	wrappedMux := manager.WrapMux(mux)
	
	initRoutes(mux, manager)
	
	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
