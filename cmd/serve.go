package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	conf := config.GetConfig()

	dbConn, err := db.NewConnectionString()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	middlewares := middleware.NewMiddlewares(conf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
