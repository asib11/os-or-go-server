package user

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if database.FindUserByEmail(newUser.Email, newUser.Password) != nil {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}

	createdUser := newUser.Store()

	utils.SendData(w, createdUser, http.StatusCreated)
}
