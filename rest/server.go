package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)
type Server struct {
	conf           *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(conf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		conf:           conf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() 
	wrappedMux := manager.WrapMux(mux)

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)


	addr := strconv.Itoa(s.conf.HttpPort)

	fmt.Println("Server is running on http://localhost:" + addr)
	err := http.ListenAndServe(":"+addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
