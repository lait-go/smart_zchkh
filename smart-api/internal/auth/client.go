package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	keycloakBaseURL = "http://localhost:9080"
	realm           = "smart-hcs"
	clientID        = "smart-hcs-client" 

	adminUsername = "admin"
	adminPassword = "admin"
	adminClientID = "admin-cli"
)

func getAdminToken() (string, error) {
	data := fmt.Sprintf(
		"client_id=%s&username=%s&password=%s&grant_type=password",
		adminClientID, adminUsername, adminPassword,
	)

	resp, err := http.Post(
		keycloakBaseURL+"/realms/master/protocol/openid-connect/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}
	if tokenResp.AccessToken == "" {
		return "", errors.New("failed to obtain access token from Keycloak")
	}

	return tokenResp.AccessToken, nil
}