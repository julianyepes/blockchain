package controllers

import (
	"fmt"
	"net/http"
	"xblock/models"
)

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blockchain: models.NewBlockchain(),
	}
}

type Blockchain struct {
	Blockchain *models.Blockchain
}

// Create is used to process the newBlockchain form when a user
// submits it. This is used to create a new blockchain.
//
// POST /create_blockchain
func (bc *Blockchain) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Invalid request")
		return
	}
	fmt.Fprintln(w, "This is a temporary response for creating a blockchain.")
}
