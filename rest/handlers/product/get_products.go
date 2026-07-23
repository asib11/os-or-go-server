package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.List()
	if err != nil {
		utils.ErrorData(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}
	utils.SendData(w, http.StatusOK, products)
}
