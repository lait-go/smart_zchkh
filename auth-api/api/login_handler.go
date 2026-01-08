package api

import (
	"auth-api/internal/auth"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	token, err := auth.LoginUserKeycloak(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials: "+err.Error(), http.StatusUnauthorized)
		return
	}

	claims, err := auth.ExtractClaimsFromToken(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to extract token claims: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":    token.AccessToken,
		"userId":   claims.Sub,
		"username": claims.PreferredUsername,
	})
}
