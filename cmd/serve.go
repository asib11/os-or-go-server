package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	conf := config.GetConfig()

	dbConn, err := db.NewConnectionString(conf.DbConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbConn, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repo
	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

	//domain
	usrSvc := user.NewService(userRepo)

	middlewares := middleware.NewMiddlewares(conf)

	productHandler := productHandler.NewHandler(middlewares, productRepo)
	userHandler := userHandler.NewHandler(conf, usrSvc)

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
