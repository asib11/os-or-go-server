package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	conf := config.GetConfig()

	// dbConn, err := db.NewConnectionString()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer dbConn.Close()

	productRepo := repo.NewProductRepo()

	middlewares := middleware.NewMiddlewares(conf)
	userRepo := repo.NewUserRepo()

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(conf, userRepo)

	server := rest.NewServer(conf, productHandler, userHandler)
	server.Start()

}
