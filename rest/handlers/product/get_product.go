package product

import (
	"net/http"
	"strconv"
	"ecommerce/utils"
)


func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	pID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorData(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.svc.Get(pID)
	if err != nil  {
		utils.ErrorData(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendData(w, http.StatusOK, product)
}