package blockchain

import (
	"encoding/hex"
)

// Blockchain đại diện cho blockchain
type Blockchain struct {
	Blocks []*Block // Danh sách các block trong blockchain
}

// AddBlock thêm một block mới vào blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	var prevBlockHash string

	if len(bc.Blocks) > 0 {
		prevBlock := bc.Blocks[len(bc.Blocks)-1]
		prevBlockHash = prevBlock.Hash
	} else {
		prevBlockHash = ""
	}

	newBlock := NewBlock(transactions, prevBlockHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// IsValid checks if the transaction is in the blockchain
func (bc *Blockchain) IsValid(transaction *Transaction) bool {
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if hex.EncodeToString(tx.Data) == hex.EncodeToString(transaction.Data) {
				return true
			}
		}
	}
	return false
}
