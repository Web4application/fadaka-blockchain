func isHashValid(hash string) bool {
    return hash[:4] == "0000"
}

func generateBlock(oldBlock Block, Data string) Block {
    var newBlock Block
    t := time.Now()

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.String()
    newBlock.Data = Data
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Nonce = 0

    for !isHashValid(newBlock.Hash) {
        newBlock.Nonce++
        newBlock.Hash = calculateHash(newBlock)
    }

    return newBlock
}
