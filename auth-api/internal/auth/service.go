package auth

import (
	"encoding/json"
	"os"
)

var userFile = "./data/user.json"

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

