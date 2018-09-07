package main

import (
	"fmt"
	"html/template"
	"net/http"

	"xblock/views"
)

var (
	homeView *views.View
	bc       *Blockchain
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, bc)
	if err != nil {
		panic(err)
	}
}

func main() {
	bc = NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	homeView = views.NewView(
		template.FuncMap{"formatShort": formatShort},
		"bulma",
		"views/home.gohtml",
		"views/layouts/block.gohtml",
	)

	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting the server on port 4000...")
	http.ListenAndServe(":4000", nil)
}
