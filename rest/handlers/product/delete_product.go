package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	pID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorData(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.productRepo.Delete(pID)
	if err != nil {
		utils.ErrorData(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	utils.SendData(w, http.StatusNoContent, "Product deleted successfully")
}
