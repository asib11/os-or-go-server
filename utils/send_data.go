package utils

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter,statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
	}
}
