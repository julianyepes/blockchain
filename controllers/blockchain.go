package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"xblock/models"
	"xblock/views"
)

func NewBlockchainC() *BlockchainC {
	return &BlockchainC{
		Blockchain: models.NewBlockchain(),
		homeView: views.NewView(
			template.FuncMap{
				"formatShort":  formatShort,
				"formatAsDate": formatAsDate,
			},
			"bulma",
			"views/home.gohtml",
			"views/components/block.gohtml",
			"views/components/newBlockForm.gohtml",
		),
	}
}

type BlockchainC struct {
	Blockchain *models.Blockchain
	homeView   *views.View
}

// Home is used to display the blockchain in the browser
//
// GET /
func (bc *BlockchainC) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	bcs := bc.Blockchain.GetBlockchainStruct()
	must(bc.homeView.Render(w, bcs))
}

// Create is used to process the newBlockchain form when a user
// submits it. This is used to create a new blockchain.
//
// POST /create_blockchain
func (bc *BlockchainC) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Invalid request")
		return
	}
	bc.Blockchain.Initialize()
	http.Redirect(w, r, "/", http.StatusFound)
}

// Add is used to process the addBlock form when a user
// submits it. This is used to add a new block to the blockchain.
//
// POST /add_block
func (bc *BlockchainC) Add(w http.ResponseWriter, r *http.Request) {
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
