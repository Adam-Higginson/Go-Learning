package main

import (
	"fmt"
	"github.com/Adam-Higginson/test-go-project/cheque"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Expected one argument of amount in pence!")
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Printf("Error when parsing amount, %v\n", err)
		return
	}

	convertedAmount, err := cheque.ConvertToChequeFormatWithDecimal(amount)
	if err != nil {
		fmt.Printf("Error when converting to cheque format, %v\n", err)
	}

	fmt.Printf("%s", convertedAmount)
}