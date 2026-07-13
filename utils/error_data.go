package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorData(w http.ResponseWriter, err string, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(err)
}
