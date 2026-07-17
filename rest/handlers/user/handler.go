package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	conf *config.Config
	userRepo repo.UserRepo
}

func NewHandler(conf *config.Config, userRepo repo.UserRepo) *Handler {
	return &Handler{
		conf:     conf,
		userRepo: userRepo,
	}
}
