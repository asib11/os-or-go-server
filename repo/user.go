package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	QUERY := `INSERT INTO users (first_name, last_name, email, password, is_shop_owner) VALUES (:first_name, :last_name, :email, :password, :is_shop_owner) RETURNING id`

	var userID int
	row, err := r.db.NamedQuery(QUERY, user)
	if err != nil {
		return nil, err
	}

	if row.Next() {
		row.Scan(&userID)
	}
	user.ID = userID
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	QUERY := `SELECT * FROM users WHERE email = $1 AND password = $2 LIMIT 1`

	var user User
	err := r.db.Get(&user, QUERY, email, password)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

