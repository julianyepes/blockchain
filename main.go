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

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/create_blockchain", bcC.Create)
	http.HandleFunc("/add_block", bcC.Add)
	fmt.Println("Starting the server on port 4000...")
	http.ListenAndServe(":4000", nil)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
