package main

import "net/http"


func createProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	receiveData(w, r, &newProduct)
	
	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)

	sendData(w, newProduct, http.StatusCreated)

}
