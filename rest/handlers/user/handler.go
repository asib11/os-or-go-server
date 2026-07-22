package user

import (
	"ecommerce/config"
)

type Handler struct {
	conf *config.Config
	svc  Service
}

func NewHandler(conf *config.Config, svc Service) *Handler {
	return &Handler{
		conf: conf,
		svc:  svc,
	}
}
