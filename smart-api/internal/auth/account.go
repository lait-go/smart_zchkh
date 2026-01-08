package auth

import (
	"encoding/json"
	"os"
)

type Account struct {
	ID         int `json:"id"`
	UserID     string `json:"user_id"`
	AccountNumber string `json:"account_number"`
	FullName   string `json:"full_name"`
	Address    string `json:"address"`
	Area       int    `json:"area"`
}

var accountFile = "./data/accounts.json"

func LoadAccounts() ([]Account, error) {
	file, err := os.Open(accountFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var accounts []Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

func SavedAccounts(accounts []Account) error {
	file, err := os.Create(accountFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(accounts)
}

func CreateAccount (newAcc Account) error {
	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	newAcc.ID = len(accounts) + 1
	accounts = append(accounts, newAcc)
	return SavedAccounts(accounts)
}