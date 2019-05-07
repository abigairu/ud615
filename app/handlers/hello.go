package handlers

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:version`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Version: "v1.0.1",
		Message: "Hello... from my fork",
	}
	json.NewEncoder(w).Encode(response)
	return
}
