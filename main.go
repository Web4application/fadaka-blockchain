package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
    Nonce     int
    Hash      string
}


var Blockchain []Block

func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Data string) Block {
    var newBlock Block
    t := time.Now()

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.Format(time.RFC3339)
    newBlock.Data = Data
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Nonce = 0

    for !isHashValid(newBlock.Hash, 4) {
        newBlock.Nonce++
        newBlock.Hash = calculateHash(newBlock)
    }

    return newBlock
}

func isHashValid(hash string, difficulty int) bool {
    prefix := ""
    for i := 0; i < difficulty; i++ {
        prefix += "0"
    }
    return hash[:difficulty] == prefix
}

func main() {
    genesisBlock := Block{0, time.Now().Format(time.RFC3339), "Genesis Block", "", "", 0}
    genesisBlock.Hash = calculateHash(genesisBlock)
    Blockchain = append(Blockchain, genesisBlock)

    newBlock := generateBlock(genesisBlock, "Web4asset Block")
    Blockchain = append(Blockchain, newBlock)

    for _, block := range Blockchain {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("PrevHash: %s\n", block.PrevHash)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Printf("Nonce: %d\n", block.Nonce)
        fmt.Println()
    }
}
