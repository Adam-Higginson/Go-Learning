package main

import (
	"encoding/json"
	"github.com/Adam-Higginson/test-go-project/cheque"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("No port specified!")
	}

	http.HandleFunc("/cheques", ChequesHandler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
type ChequeConversionRequest struct {
	AmountInPence int `json:"amountInPence"`
}

type ChequeResponse struct {
	AmountInPence int `json:"amountInPence"`
	ChequeString string `json:"chequeString"`
}

func ChequesHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var chequeConversionRequest ChequeConversionRequest
	err := json.Unmarshal(body, &chequeConversionRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	amountInPence := chequeConversionRequest.AmountInPence
	chequeString, err := cheque.ConvertToChequeFormat(amountInPence)
	if err != nil {
		log.Printf("Error when attempting to convert cheque amount, %v", err)
		http.Error(w, "Error when attempting to convert cheque amount", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	chequeResponse := ChequeResponse{
		AmountInPence: amountInPence,
		ChequeString:  chequeString,
	}

	json.NewEncoder(w).Encode(chequeResponse)
}