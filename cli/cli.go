package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Blockchain-va-ng-d-ng-20MMT-HCMUS-2024/lab1/blockchain"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bc := &blockchain.Blockchain{}

	for {
		fmt.Print("Enter command (addblock, checkvalid, print): ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "addblock":
			var transactions []*blockchain.Transaction
			for {
				fmt.Print("Enter transaction data (or 'done' to finish): ")
				scanner.Scan()
				data := scanner.Text()
				if data == "done" {
					break
				}
				tx := &blockchain.Transaction{Data: []byte(data)}
				transactions = append(transactions, tx)
			}
			if len(transactions) > 0 {
				bc.AddBlock(transactions)
				fmt.Println("Block added with transactions.")
			} else {
				fmt.Println("No transactions entered. Block not added.")
			}
		case "checkvalid":
			fmt.Print("Enter transaction data: ")
			scanner.Scan()
			data := scanner.Text()
			tx := &blockchain.Transaction{Data: []byte(data)}
			if bc.IsValid(tx) {
				fmt.Println("Transaction is valid.")
			} else {
				fmt.Println("Transaction is invalid.")
			}
		case "print":
			for _, block := range bc.Blocks {
				fmt.Printf("Prev. hash: %s\n", block.PrevBlockHash)
				fmt.Printf("Hash: %s\n", block.Hash)
				fmt.Printf("Merkle Root: %s\n", block.MerkleRoot)
				fmt.Printf("Transactions:\n")
				for _, tx := range block.Transactions {
					fmt.Printf("- Tx: %s\n", string(tx.Data))
				}
				fmt.Println()
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}
