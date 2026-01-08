package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"smart-api/internal/httpx"
	"strings"
)

const chargeFile = "./data/charges.json"

var ErrValidation = errors.New("validation failed")
var ErrNotFound = errors.New("resource not found")

func LoadCharges() ([]Charge, error) {
	file, err := os.Open(chargeFile)
	if err != nil {
		return []Charge{}, err
	}
	defer file.Close()

	var charges []Charge
	err = json.NewDecoder(file).Decode(&charges)
	if err != nil {
		return []Charge{}, err
	}

	return charges, err
}

func SaveCharges(charges []Charge) error {
	file, err := os.Create(chargeFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(charges)
}

func CreateCharge(newCharge Charge) error {
	charges, err := LoadCharges()
	if err != nil {
		return err
	}

	if err := newCharge.Validate(); err != nil {
		return err
	}

	newCharge.ID = len(charges) + 1
	charges = append(charges, newCharge)
	return SaveCharges(charges)
}

func (c *Charge) Validate() error {
	if len(c.UserId) <= 0 {
		return fmt.Errorf("invalid user id: %w", ErrValidation)
	}
	if c.Amount <= 0 {
		return fmt.Errorf("invalid amount: %w", ErrValidation)
	}
	if strings.TrimSpace(c.Category) == "" {
		return fmt.Errorf("invalid category: %w", ErrValidation)
	}
	if strings.TrimSpace(c.Date) == "" {
		return fmt.Errorf("invalid date: %w", ErrValidation)
	}
	return nil
}

func PaydingFunc(ch Charge) error {
	us, _ := LoadUser()

	var anem string
	for _, elem := range us{
		if elem.ID == ch.UserId{
			anem = elem.Username
		}
	}
	Paid := Paidin{
		Name: anem,
		User_id: ch.UserId,
		Data: ch.Date,
		Category: ch.Category,
		Amount: ch.Amount,
		Qr_data: "https://example.com/payment?invoice=INV-TEST-2025-06-003&amount=13000.75",
	}

	jsonData, err := json.Marshal(Paid)
	if err != nil{
		return err
	}

	fmt.Println(Paid)
	_, err = http.Post("http://localhost:8001/generate-receipt", "application/json", bytes.NewBuffer(jsonData))
	if err != nil{
		return err
	}
	return nil
}

func SendValidationError(w http.ResponseWriter, err error) {
	httpx.NewJSONError(w, http.StatusBadRequest, "Validation failed", err.Error())
}

func IsValidationError(err error) bool {
	return errors.Is(err, ErrValidation)
}
