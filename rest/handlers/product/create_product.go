package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.ErrorData(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	createdProduct, err := h.svc.Create(domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
	})

	if err != nil {
		utils.ErrorData(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.SendData(w, http.StatusCreated, createdProduct)
}
