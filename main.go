package main

import (
	"fmt"
	"net/http"
	"os"

	"xblock/controllers"
)

func main() {

	bcC := controllers.NewBlockchainC()
	defer bcC.Blockchain.CloseDB()

	http.HandleFunc("/", bcC.Home)
	http.HandleFunc("/create_blockchain", bcC.Create)
	http.HandleFunc("/add_block", bcC.Add)
	fmt.Println("Starting the server on port 4000...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
