package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"math/big"
	"bytes"
)

// Define Fadaka blockchain API endpoint (this is a placeholder)
const fadakaAPIUrl = "http://localhost:4000/api/transactions" // Example API endpoint for Fadaka blockchain

// Structure to hold wallet details
type Wallet struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// Function to generate a new wallet (Private & Public key pair)
func generateWallet() (*Wallet, error) {
	// Generate elliptic curve private key
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("error generating key: %v", err)
	}

	// Derive the public key from the private key
	pub := &priv.PublicKey

	// Create the wallet with public and private keys in hexadecimal format
	privateKeyHex := hex.EncodeToString(priv.D.Bytes())
	publicKeyHex := fmt.Sprintf("%x", pub.X) + fmt.Sprintf("%x", pub.Y)

	return &Wallet{
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}, nil
}

// Function to send the generated wallet to Fadaka blockchain (via an API call)
func storeWalletInFadaka(wallet *Wallet) error {
	// Construct the payload to send to Fadaka blockchain
	payload := map[string]string{
		"public_key":  wallet.PublicKey,
		"private_key": wallet.PrivateKey,
	}

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %v", err)
	}

	// Make a POST request to the Fadaka blockchain to store the wallet
	resp, err := http.Post(fadakaAPIUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending request to Fadaka: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to store wallet in Fadaka, status: %v", resp.StatusCode)
	}

	// Assuming Fadaka blockchain returns a success response
	log.Println("Wallet successfully stored in Fadaka blockchain.")
	return nil
}

// HTTP handler to generate wallet and store it in Fadaka blockchain
func generateWalletHandler(w http.ResponseWriter, r *http.Request) {
	// Generate a new wallet (private & public keys)
	wallet, err := generateWallet()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating wallet: %v", err), http.StatusInternalServerError)
		return
	}

	// Store the wallet in the Fadaka blockchain
	err = storeWalletInFadaka(wallet)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error storing wallet in Fadaka: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the generated wallet details as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallet)
}

func main() {
	// Route to generate and store wallet
	http.HandleFunc("/generate_wallet", generateWalletHandler)

	// Start the server
	port := ":5000" // Default port for the backend
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
