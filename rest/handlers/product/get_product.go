package product

import (
	"net/http"
	"strconv"
	"ecommerce/database"
	"ecommerce/utils"
)


func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	pID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.GetProductByID(pID)
	if product == nil {
		utils.ErrorData(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}