package main

import (
	"encoding/json"
	"net/http"
)

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
	}
}
