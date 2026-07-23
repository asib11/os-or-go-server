package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	pID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorData(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var newProduct ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product, err := h.svc.Update(domain.Product{
		ID:          pID,
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImageURL:    newProduct.ImageURL,
	})

	if err != nil {
		utils.ErrorData(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.SendData(w, http.StatusOK, product)
}
