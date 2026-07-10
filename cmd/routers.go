package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /test",
		manager.With(
			http.HandlerFunc(handlers.TestHandler),
		))

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		))

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProducts),
		))

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		))
}
