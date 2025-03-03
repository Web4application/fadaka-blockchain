package Wallet_Management

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "strconv"
)

// Wallet represents a cryptocurrency wallet with a private and public key.
type Wallet struct {
    PrivateKey *ecdsa.PrivateKey
    PublicKey  string
}

// createWallet generates a new Wallet with an ECDSA private and public key pair.
func createWallet() (Wallet, error) {
    privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return Wallet{}, fmt.Errorf("failed to generate private key: %v", err)
    }
    pubKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
    return Wallet{
        PrivateKey: privKey,
        PublicKey:  hex.EncodeToString(pubKey),
    }, nil
}

// signTransaction signs a transaction using the private key and returns the signature.
func signTransaction(tx Transaction, privKey *ecdsa.PrivateKey) (string, error) {
    hash := sha256.Sum256([]byte(tx.Sender + tx.Receiver + strconv.Itoa(tx.Amount) + tx.Timestamp))
    r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
    if err != nil {
        return "", fmt.Errorf("failed to sign transaction: %v", err)
    }
    return hex.EncodeToString(r.Bytes()) + hex.EncodeToString(s.Bytes()), nil
}

// Transaction represents a cryptocurrency transaction.
type Transaction struct {
    Sender    string
    Receiver  string
    Amount    int
    Timestamp string
}

func main() {
    wallet, err := createWallet()
    if err != nil {
        log.Fatalf("Error creating wallet: %v", err)
    }
    fmt.Printf("Wallet created: %+v\n", wallet)
}
