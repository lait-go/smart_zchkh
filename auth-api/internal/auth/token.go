package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType   string `json:"token_type"`
}

func LoginUserKeycloak(username, password string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)

	resp, err := http.Post(
		keycloakBaseURL+"/realms/"+realm+"/protocol/openid-connect/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()
	fmt.Println("Trying to login:", username)
	fmt.Println("POST:", data.Encode())
	fmt.Println("Resp.Status:", resp.StatusCode)
	if resp.StatusCode != 200 {
		var bodyBytes bytes.Buffer
		bodyBytes.ReadFrom(resp.Body)
		return nil, errors.New("login failed: " + bodyBytes.String())
	}
	fmt.Println("Trying to login:", username)
	fmt.Println("POST:", data.Encode())
	fmt.Println("Resp.Status:", resp.StatusCode)
	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	fmt.Println("Trying to login:", username)
	fmt.Println("POST:", data.Encode())
	fmt.Println("Resp.Status:", resp.StatusCode)

	return &tokenResp, nil
}


type TokenClaims struct {
	Sub               string `json:"sub"`
	PreferredUsername string `json:"preferred_username"`
}

func ExtractClaimsFromToken(token string) (TokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return TokenClaims{}, fmt.Errorf("invalid token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return TokenClaims{}, fmt.Errorf("base64 decode failed: %w", err)
	}

	var claims TokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return TokenClaims{}, fmt.Errorf("json parse failed: %w", err)
	}

	return claims, nil
}
