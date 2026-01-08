package handlers

const filePath = "../../smart-api/data/charges.json"

const accountsFilePath = "../../smart-api/data/accounts.json"

type Account struct {
	ID            uint   `json:"id"`
	UserID        string `json:"user_id"`
	AccountNumber string `json:"account_number"`
	FullName      string `json:"full_name"`
	Address       string `json:"address"`
	Area          int    `json:"area"`
}

type Transaction struct {
	ID       uint    `json:"id"`
	UserID   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Date     string  `json:"date"`
	Category string  `json:"category"`
	Paid     bool    `json:"paid"`
}

type TransactionResponse struct {
	Transaction
	UserName  string `json:"user_name"`
	AccountNumber string `json:"account_number"`
}
