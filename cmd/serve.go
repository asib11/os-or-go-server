package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {
	conf := config.GetConfig()
	rest.Start(conf)

}
