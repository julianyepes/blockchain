package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"xblock/models"
)

func NewBlockchain() *Blockchain {
	/* return &Blockchain{
		Blockchain: models.NewBlockchain(),
	} */
	return &Blockchain{
		Blockchain: &models.Blockchain{
			Blocks: nil,
		},
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
	bc.Blockchain = models.NewBlockchain()
	http.Redirect(w, r, "/", http.StatusFound)
}

// Add is used to process the addBlock form when a user
// submits it. This is used to add a new block to the blockchain.
//
// POST /add_block
func (bc *Blockchain) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Invalid request")
		return
	}
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	data := strings.Join(r.PostForm["data"], " ")
	bc.Blockchain.AddBlock(data)
	http.Redirect(w, r, "/", http.StatusFound)
}
