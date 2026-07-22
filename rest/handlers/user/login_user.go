package user

import (
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.ErrorData(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		utils.ErrorData(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if usr == nil {
		utils.ErrorData(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	accessToken, err := utils.CreateJWT(h.conf.JWTSecret, utils.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,


	})
	if err != nil {
		utils.ErrorData(w, http.StatusInternalServerError, "Failed to create access token")
		return
	}

	utils.SendData(w, http.StatusOK, accessToken)
}
