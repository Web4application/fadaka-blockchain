import std/[sha1, times, strutils]
import nimcrypto # Standard for blockchain hashing

type
  Block = object
    index: int
    timestamp: float
    data: string
    previousHash: string
    hash: string

proc calculateHash(index: int, timestamp: float, data: string, prevHash: string): string =
  # Uses nimcrypto or standard sha
  let payload = $index & $timestamp & data & prevHash
  return $sha1.compute(payload)

proc createGenesisBlock(): Block =
  let ts = epochTime()
  let h = calculateHash(0, ts, "Genesis Block", "0")
  result = Block(index: 0, timestamp: ts, data: "Genesis Block", previousHash: "0", hash: h)

proc createNewBlock(prevBlock: Block, data: string): Block =
  let index = prevBlock.index + 1
  let ts = epochTime()
  let h = calculateHash(index, ts, data, prevBlock.hash)
  result = Block(index: index, timestamp: ts, data: data, previousHash: prevBlock.hash, hash: h)

# Main Execution Loop
when isMainModule:
  var blockchain = @[createGenesisBlock()]
  echo "Fadaka Genesis Block: ", blockchain[0].hash

  # Add a test block
  let nextBlock = createNewBlock(blockchain[0], "Fadaka Transaction Data")
  blockchain.add(nextBlock)
  echo "Next Block Hash: ", nextBlock.hash
