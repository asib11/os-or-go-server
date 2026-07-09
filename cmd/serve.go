package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Test, middleware.Logger) 

	mux := http.NewServeMux() // Create a new ServeMux to handle routes. its called router in other languages

	initRoutes(mux, manager)

	
	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
