package handlers

import (
	"ecommerce/utils"
	"ecommerce/database"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	pID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	database.DeleteProduct(pID)

	utils.SendData(w, "Product deleted successfully", 204)
}
