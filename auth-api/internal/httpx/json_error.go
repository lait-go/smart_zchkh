package httpx

import (
	"encoding/json"
	"net/http"
)

type JSONERROR struct {
	Error string `json:"error"`
	Details   string `json:"detials,omitempty"`
}

func NewJSONError(w http.ResponseWriter, status int, message string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	jsonErr := JSONERROR{
		Error: message,
		Details: details ,
	}
	
	json.NewEncoder(w).Encode(jsonErr)
}