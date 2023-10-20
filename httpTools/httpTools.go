package httptools

import (
	"encoding/json"
	"net/http"
)

type SimpleRequest struct {
	Message string `json:"message"`
}

type Request struct {
	ID    string
	Key   string
	Value string
}
type Response struct {
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
