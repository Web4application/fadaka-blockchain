package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
)

// Open SQLite DB (create if not exists)
var db, _ = sql.Open("sqlite3", "./fadaka.db")

func init() {
	// Create table if not exists
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS accounts (address TEXT PRIMARY KEY, balance INTEGER)")
	if err != nil {
		fmt.Println("Error creating table:", err)
	}
}

// Get wallet balance
func getBalance(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")

	var balance int
	err := db.QueryRow("SELECT balance FROM accounts WHERE address = ?", address).Scan(&balance)
	if err != nil {
		http.Error(w, "Address not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"address": address,
		"balance": balance,
	}
	json.NewEncoder(w).Encode(response)
}

// Send transaction (simplified)
func sendTransaction(w http.ResponseWriter, r *http.Request) {
	var txData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&txData); err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	from := txData["from"]
	to := txData["to"]
	value := 1 // Simplified value

	// Update balances in DB
	tx, _ := db.Begin()
	defer tx.Commit()

	_, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE address = ?", value, from)
	if err != nil {
		http.Error(w, "Error updating sender balance", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE address = ?", value, to)
	if err != nil {
		http.Error(w, "Error updating recipient balance", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"from": from,
		"to": to,
		"value": value,
		"status": "success",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/rpc/fadaka_getBalance", getBalance)
	http.HandleFunc("/rpc/fadaka_sendTransaction", sendTransaction)

	fmt.Println("Fadaka blockchain RPC server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
