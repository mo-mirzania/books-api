package utils

import (
	"encoding/json"
	"net/http"
)

// Ping func
func Ping(w http.ResponseWriter, r *http.Request) {
	var u = struct {
		Ping string `json:"ping"`
	}{
		Ping: "pong",
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&u)
}
