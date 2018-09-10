package main

import (
	"fmt"
	"html/template"
	"net/http"

	"xblock/controllers"
	"xblock/views"
)

var (
	homeView *views.View
	bcC      *controllers.Blockchain
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, bcC.Blockchain))
}

func main() {
	homeView = views.NewView(
		template.FuncMap{"formatShort": formatShort},
		"bulma",
		"views/home.gohtml",
		"views/components/block.gohtml",
	)
	bcC = controllers.NewBlockchain()

	bcC.Blockchain.AddBlock("Send 1 BTC to Ivan")
	bcC.Blockchain.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bcC.Blockchain.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting the server on port 4000...")
	http.ListenAndServe(":4000", nil)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
