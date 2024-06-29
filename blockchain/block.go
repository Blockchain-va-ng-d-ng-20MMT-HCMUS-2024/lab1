package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64          // Thời điểm tạo block
	Transactions  []*Transaction // Danh sách các giao dịch trong block
	PrevBlockHash string         // Hash của block trước đó
	Hash          string         // Hash của block hiện tại
	MerkleRoot    string         // Gốc Merkle của các giao dịch
	Nonce         int            // Nonce để Proof of Work
}

// CalculateHash tính toán hash của block
func (b *Block) CalculateHash() {
	headers := b.PrevBlockHash +
		strconv.FormatInt(b.Timestamp, 10) +
		b.MerkleRoot +
		strconv.Itoa(b.Nonce)

	hashInBytes := sha256.Sum256([]byte(headers))
	b.Hash = hex.EncodeToString(hashInBytes[:])
}

// NewBlock tạo một block mới
func NewBlock(transactions []*Transaction, prevBlockHash string) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    "",
		Nonce:         0,
	}
	block.MerkleRoot = block.CalculateMerkleRoot()
	block.CalculateHash()
	return block
}
