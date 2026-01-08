package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"charge_service/internal/httpx"
	"charge_service/internal/logger"
)

var log logger.Log = &logger.Console{}

func readTransactions() ([]Transaction, error) {
	var txs []Transaction

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return txs, nil
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &txs)
	return txs, err
}

func readAccounts() ([]Account, error) {
	var accs []Account

	if _, err := os.Stat(accountsFilePath); os.IsNotExist(err) {
		return accs, nil
	}

	data, err := ioutil.ReadFile(accountsFilePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &accs)
	return accs, err
}

func writeTransactions(txs []Transaction) error {
	data, err := json.MarshalIndent(txs, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func CreateTransaction(c *gin.Context) {
	var tx Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		log.Errorf("Invalid JSON: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	txs, err := readTransactions()
	if err != nil {
		log.Errorf("Read error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to read file", err.Error())
		return
	}

	var maxID uint = 0
	for _, existingTx := range txs {
		if existingTx.ID > maxID {
			maxID = existingTx.ID
		}
	}
	tx.ID = maxID + 1

	txs = append(txs, tx)

	if err := writeTransactions(txs); err != nil {
		log.Errorf("Write error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to write data", err.Error())
		return
	}

	log.Info("Transaction created: %+v", tx)
	c.JSON(http.StatusOK, tx)
}

func GetTransactions(c *gin.Context) {
	txs, err := readTransactions()
	if err != nil {
		log.Errorf("Read error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to read transactions", err.Error())
		return
	}

	accounts, err := readAccounts()
	if err != nil {
		log.Errorf("Accounts read error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to read accounts", err.Error())
		return
	}

	accountMap := make(map[string]Account)
	for _, acc := range accounts {
		accountMap[acc.UserID] = acc
	}

	var result []TransactionResponse
	for _, tx := range txs {
		match := true

		if v := c.Query("id"); v != "" {
			id, _ := strconv.Atoi(v)
			if tx.ID != uint(id) {
				match = false
			}
		}
		
		if v := c.Query("user_id"); v != "" && tx.UserID != v {
			match = false
		}
		if v := c.Query("amount"); v != "" {
			a, _ := strconv.ParseFloat(v, 64)
			if tx.Amount != a {
				match = false
			}
		}
		if v := c.Query("date"); v != "" && tx.Date != v {
			match = false
		}
		if v := c.Query("category"); v != "" && tx.Category != v {
			match = false
		}
		if v := c.Query("paid"); v != "" {
			p, _ := strconv.ParseBool(v)
			if tx.Paid != p {
				match = false
			}
		}

		acc := accountMap[tx.UserID]
		if v := c.Query("account_number"); v != "" && acc.AccountNumber != v {
			match = false
		}
		if v := c.Query("full_name"); v != "" && acc.FullName != v {
			match = false
		}
		if v := c.Query("address"); v != "" && acc.Address != v {
			match = false
		}
		if v := c.Query("area"); v != "" {
			a, _ := strconv.Atoi(v)
			if acc.Area != a {
				match = false
			}
		}

		if match {
			result = append(result, TransactionResponse{
				Transaction:   tx,
				UserName: acc.FullName,
				AccountNumber: acc.AccountNumber,
			})
		}
	}

	log.Info("Returned %d filtered transactions", len(result))
	c.JSON(http.StatusOK, result)
}


func PutTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.NewJSONError(c.Writer, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	var updatedTx Transaction
	if err := c.ShouldBindJSON(&updatedTx); err != nil {
		httpx.NewJSONError(c.Writer, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	txs, err := readTransactions()
	if err != nil {
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to read file", err.Error())
		return
	}

	found := false
	for i, tx := range txs {
		if tx.ID == uint(id) {
			updatedTx.ID = tx.ID
			txs[i] = updatedTx
			found = true
			break
		}
	}

	if !found {
		httpx.NewJSONError(c.Writer, http.StatusNotFound, "Transaction not found", "")
		return
	}

	if err := writeTransactions(txs); err != nil {
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to write data", err.Error())
		return
	}

	log.Info("Transaction updated: %+v", updatedTx)
	c.JSON(http.StatusOK, updatedTx)
}
