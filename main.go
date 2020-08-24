package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("No port specified!")
	}

	http.HandleFunc("/cheques/", ChequesHandler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func ChequesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}