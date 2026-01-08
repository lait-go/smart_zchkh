package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"smart-api/internal/auth"
	"smart-api/internal/httpx"
	"strconv"
	"strings"
)

func ChargesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		httpx.NewJSONError(w, http.StatusBadRequest, "Missing user_id", "user_id is required")
		return
	}

	charges, err := auth.LoadCharges()
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load charges", err.Error())
		return
	}

	var userCharges []auth.Charge
	for _, c := range charges {
		if c.UserId == userIDStr {
			userCharges = append(userCharges, c)
		}
	}

	json.NewEncoder(w).Encode(userCharges)
}

func CreateChargeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	

	var newCharge auth.Charge
	if err := json.NewDecoder(r.Body).Decode(&newCharge); err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid request", "Invalid request")
		return
	}

	if err := newCharge.Validate(); err != nil {
	auth.SendValidationError(w, err)
	return
}

	charges, err := auth.LoadCharges()
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load charges", err.Error())
		return
	}

	newCharge.ID = len(charges) + 1
	charges = append(charges, newCharge)

	err = auth.SaveCharges(charges)
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to save charges", err.Error())
		return
	}

	json.NewEncoder(w).Encode(newCharge)
}

func ChargeHandlerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET allowed")
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 || parts[3] == "" {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid path", "Expected /api/charges/{id}")
		return
	}

	id, err := strconv.Atoi(parts[3])
	if err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid id", "Must be number")
		return
	}

	charge, err := GetCharge(id)
	if err != nil {
		if errors.Is(err, auth.ErrNotFound) {
			httpx.NewJSONError(w, http.StatusNotFound, "Charge not found", err.Error())
			return
		}
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to get charge", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charge)
}

func GetCharge(id int) (auth.Charge, error) {
	charges, err := auth.LoadCharges()
	if err != nil {
		return auth.Charge{}, err
	}

	for _, charge := range charges {
		if charge.ID == id {
			return charge, nil
		}
	}
	return auth.Charge{}, fmt.Errorf("charge not found: %w", auth.ErrNotFound)
}

func ChargesHandlerPut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only PUT allowed")
		return
	}

	charges, err := auth.LoadCharges()
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load charges", err.Error())
		return
	}

	var updatedCharge auth.Charge
	if err := json.NewDecoder(r.Body).Decode(&updatedCharge); err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid request", "Invalid JSON body")
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid path", "Expected /api/charges/{id}")
		return
	}

	idStr := parts[3]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid id", "ID must be integer")
		return
	}

	updated := false

	for i, c := range charges {
		if c.ID == id {
			charges[i].Paid = true
			updatedCharge = charges[i]
			updated = true
			break
		}
	}

	if !updated {
		httpx.NewJSONError(w, http.StatusNotFound, "Charge not found", "No matching charge found")
		return
	}

	if err := auth.SaveCharges(charges); err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to save charges", err.Error())
		return
	}

	auth.PaydingFunc(updatedCharge)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCharge)
}
