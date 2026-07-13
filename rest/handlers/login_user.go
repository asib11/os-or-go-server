package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	usr := database.FindUserByEmail(loginReq.Email, loginReq.Password)
	if usr == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	utils.SendData(w, usr, http.StatusOK)
}
