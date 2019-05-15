package handlers

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Version: "v1.0.3",
		Message: "Hello... world?",
	}
	json.NewEncoder(w).Encode(response)
	return
}
