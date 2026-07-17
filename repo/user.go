package repo

import "errors"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	return repo
}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
