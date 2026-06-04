package handlers

import (
	"ecommerce/product"
	"ecommerce/utils"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct product.Product

	utils.ReceiveData(w, r, &newProduct)

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	utils.SendData(w, newProduct, http.StatusCreated)
}
