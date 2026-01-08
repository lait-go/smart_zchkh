package auth

type User struct {
	ID       string
	Username string
	Password string
}

type Charge struct {
	ID       int    `json:"id"`
	UserId   string `json:"user_id"`
	Amount   int    `json:"amount"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Paid     bool   `json:"paid"`
}

type Paidin struct {
	Name     string `json:"payer_name"`
	User_id  string `json:"invoice_number"`
	Data     string `json:"invoice_date"`
	Category string `json:"services"`
	Amount   int    `json:"total_amount"`
	Qr_data  string `json:"qr_data"`
}
