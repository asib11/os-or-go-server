package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "Failed to list products", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, http.StatusOK, products)
}
