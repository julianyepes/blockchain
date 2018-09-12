package main

import (
	"fmt"
	"net/http"

	"xblock/controllers"
)

func main() {

	bcC := controllers.NewBlockchainC()
	defer bcC.Blockchain.CloseDB()

	http.HandleFunc("/", bcC.Home)
	http.HandleFunc("/create_blockchain", bcC.Create)
	http.HandleFunc("/add_block", bcC.Add)
	fmt.Println("Starting the server on port 4000...")
	http.ListenAndServe(":4000", nil)
}
