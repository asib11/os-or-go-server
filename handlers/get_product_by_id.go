package handlers

import (
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/utils"
)


func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}