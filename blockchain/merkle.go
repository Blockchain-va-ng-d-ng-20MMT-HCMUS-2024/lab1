package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleNode đại diện cho một node trong cây Merkle
type MerkleNode struct {
	Left  *MerkleNode // Node con trái
	Right *MerkleNode // Node con phải
	Data  string      // Dữ liệu của node
}

// MerkleTree đại diện cho một cây Merkle
type MerkleTree struct {
	Root *MerkleNode // Gốc của cây Merkle
}

// NewMerkleNode tạo một node Merkle mới
func NewMerkleNode(left, right *MerkleNode, data string) *MerkleNode {
	node := &MerkleNode{}
	if left == nil && right == nil {
		node.Data = data
	} else {
		var prevHash string
		if right == nil {
			prevHash = left.Data + left.Data
		} else {
			prevHash = left.Data + right.Data
		}
		hashInBytes := sha256.Sum256([]byte(prevHash))
		node.Data = hex.EncodeToString(hashInBytes[:])
		node.Left = left
		node.Right = right
	}
	return node
}

// NewMerkleTree tạo một cây Merkle từ các giao dịch
func NewMerkleTree(transactions []*Transaction) *MerkleTree {
	var nodes []*MerkleNode

	// Tạo các node lá từ các giao dịch
	for _, tx := range transactions {
		node := NewMerkleNode(nil, nil, string(tx.Data))
		nodes = append(nodes, node)
	}

	// Xây dựng cây Merkle từ các node lá
	for len(nodes) > 1 {
		var newLevel []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			var right *MerkleNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			} else {
				right = left
			}

			node := NewMerkleNode(left, right, "")
			newLevel = append(newLevel, node)
		}

		nodes = newLevel
	}

	tree := &MerkleTree{Root: nodes[0]}
	return tree
}

// CalculateMerkleRoot tính toán gốc Merkle của cây Merkle
func (b *Block) CalculateMerkleRoot() string {
	merkleTree := NewMerkleTree(b.Transactions)
	return merkleTree.Root.Data
}
