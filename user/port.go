package user

import (
	"ecommerce/domain"
	usrHandlers "ecommerce/rest/handlers/user"
)



type Service interface {
	usrHandlers.Service // embedding -> embedding the Service interface from the user handler package
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
	// List() ([]*domain.User, error)
	// Update(user domain.User) (*domain.User, error)
	// Delete(userID int) error
}