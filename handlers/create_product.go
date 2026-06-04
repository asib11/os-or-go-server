package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	utils.ReceiveData(w, r, &newProduct)

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, http.StatusCreated)
}
