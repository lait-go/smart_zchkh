package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
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

func CreateKeycloakUser(username, password, email, first_name, last_name string) error {
	token, err := getAdminToken()
	if err != nil {
		return fmt.Errorf("cannot get admin token: %w", err)
	}

	userData := map[string]interface{}{
	"username":       username,
	"email":          email,
	"firstName":      first_name,
	"lastName":       last_name,
	"enabled":        true,
	"emailVerified":  true,
}

	body, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/admin/realms/%s/users", keycloakBaseURL, realm), bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("create user request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes := new(bytes.Buffer)
		bodyBytes.ReadFrom(resp.Body)
		return fmt.Errorf("failed to create user: %s", bodyBytes.String())
	}

	location := resp.Header.Get("Location")
	if location == "" {
		return errors.New("user created but no Location header returned")
	}
	parts := strings.Split(location, "/")
	userID := parts[len(parts)-1]

	passData := map[string]interface{}{
		"type":      "password",
		"value":     password,
		"temporary": false,
	}
	passBody, _ := json.Marshal(passData)
	passReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/admin/realms/%s/users/%s/reset-password", keycloakBaseURL, realm, userID), bytes.NewBuffer(passBody))
	passReq.Header.Set("Authorization", "Bearer "+token)
	passReq.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(passReq)
	if err != nil {
		return fmt.Errorf("set password request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes := new(bytes.Buffer)
		bodyBytes.ReadFrom(resp.Body)
		return fmt.Errorf("failed to set password: %s", bodyBytes.String())
	}

	users, err := LoadUser()
	if err != nil {
		return err
	}

	user := User{
		ID: userID,
		Username: username,
	}

	users = append(users, user)
	SaveUsers(users)

	return nil
}


