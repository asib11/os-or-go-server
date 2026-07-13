package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(conf config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux() // Create a new ServeMux to handle routes. its called router in other languages
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := strconv.Itoa(conf.HttpPort)

	fmt.Println("Server is running on http://localhost:" + addr)
	err := http.ListenAndServe(":"+addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
