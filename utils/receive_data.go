package utils

import (
	"encoding/json"
	"net/http"
)

func ReceiveData(w http.ResponseWriter, r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		http.Error(w, "Error decoding data", http.StatusBadRequest)
		return err
	}
	return nil
}
