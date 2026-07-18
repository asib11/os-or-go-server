package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
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

	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

	middlewares := middleware.NewMiddlewares(conf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(conf, userRepo)

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
