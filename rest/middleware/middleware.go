package middleware

import "ecommerce/config"

type Middlewares struct {
	conf *config.Config
}

func NewMiddlewares(conf *config.Config) *Middlewares {
	return &Middlewares{
		conf: conf,
	}
}