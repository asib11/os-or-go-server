package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorData(w http.ResponseWriter,statusCode int, err string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(err)
}
