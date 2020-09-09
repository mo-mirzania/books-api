package utils

import (
	"encoding/json"
	"net/http"

	"github.com/mo-mirzania/api/books-list/model"
)

// SendError func
func SendError(w http.ResponseWriter, status int, err model.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// SendSuccess func
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
