package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const userFile = "./data/user.json"
var keycloakURL = "http://localhost:8081"


func LoadUser() ([]User, error) {
	file, err := os.Open(userFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUsers(users []User) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}

func RegisterUser(username, password string) error {
	token, err := getAdminToken()
	if err != nil {
		return fmt.Errorf("failed to get admin token: %w", err)
	}

	userData := map[string]interface{}{
		"username": username,
		"enabled":  true,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     password,
				"temporary": false,
			},
		},
	}

	body, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %w", err)
	}

	req, err := http.NewRequest("POST", keycloakURL+"/admin/realms/smart-hcs/users", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("keycloak error (%d): %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func LoginUser(username, password string) (*User, error) {
	users, err := LoadUser()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}

	return nil, errors.New("invalid credentials")
}

func (u *User) Validate() error {
	if len(u.Username) < 2 {
	fmt.Println(2)
		return fmt.Errorf("username must be at least 2 characters: %w", ErrValidation)
	}
	// if  len(u.Password) < 4 {
	// fmt.Println(3)
	// 	return fmt.Errorf("password must be at least 6 characters: %w", ErrValidation)
	// }
	if u.Username == "" {
	fmt.Println(4)
		return fmt.Errorf("invalid username: %w", ErrValidation)
	}
	if u.Password == "" {
	fmt.Println(5)
		return fmt.Errorf("invalid password: %w", ErrValidation)
	}
	return nil
}