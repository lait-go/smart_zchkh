package api

import (
	"encoding/json"
	"net/http"
	"smart-api/internal/auth"
	"smart-api/internal/httpx"
	"strings"
)

func AccountsHadlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	userIDParam := r.URL.Query().Get("user_id")

	accounts, err := auth.LoadAccounts()
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load accounts", err.Error())
		return
	}

	if userIDParam == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)
		return
	}

	var filtered []auth.Account
	for _, acc := range accounts {
		if acc.UserID == userIDParam {
			filtered = append(filtered, acc)
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}


func AccountsHadlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}
	
	var newAcc auth.Account
	if err := json.NewDecoder(r.Body).Decode(&newAcc); err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid request", "Invalid request")
		return
	}

	err := auth.CreateAccount(newAcc)
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to create account", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Account created successfully"))
}

func AccountsHandlerByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET allowed")
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid path", "Expected /api/accounts/{user_id}")
		return
	}

	userID := parts[3]

	result, err := GetAccounts(userID)
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to get accounts", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetAccounts(id string) ([]auth.Account, error) {
	accounts, err := auth.LoadAccounts()
	if err != nil {
		return nil, err
	}

	var results []auth.Account
	for _, acc := range accounts {
		if acc.UserID == id {
			results = append(results, acc)
		}
	}

	return results, nil
}