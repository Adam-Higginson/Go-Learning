package cheque

import (
	"errors"
	"fmt"
	"math"
)

var belowTwenty = [20]string{
	"Zero",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
}

var belowHundred = map[int]string {
	2 : "Twenty",
	3 : "Thirty",
	4 : "Forty",
	5 : "Fifty",
	6 : "Sixty",
	7 : "Seventy",
	8 : "Eighty",
	9 : "Ninety",
}

func ConvertToChequeFormat(amountInPence int) (string, error) {
	if amountInPence < 0 {
		return "", errors.New("amount in pence is less than zero")
	}

	poundsAmount := amountInPence / 100
	penceAmount := amountInPence % 100
	return fmt.Sprintf("%s%s And %sPence",
		convertToChequeFormatRecurse("", poundsAmount),
		getPoundsString(poundsAmount),
		convertToChequeFormatRecurse("", penceAmount)), nil
}

func convertToChequeFormatRecurse(currentString string, amount int) string {
	switch {
	case amount < 20:
		return calculatePoundUnits(currentString, amount)
	case amount < 100:
		return calculateTens(currentString, amount)
	default:
		return calculateHundredsAndGreater(currentString, amount)
	}
}

func calculateTens(currentString string, amount int) string {
	currentDigit := amount / 10
	convertedToString := currentString + belowHundred[currentDigit] + " "
	return convertToChequeFormatRecurse(convertedToString, amount% 10)
}

func calculateHundredsAndGreater(currentString string, amount int) string {
	divider, name := getFactorForAmount(amount)

	currentDigit := amount / divider
	prefix := convertToChequeFormatRecurse("", currentDigit)
	convertedToString := currentString + prefix + name + " "

	return convertToChequeFormatRecurse(convertedToString, amount % divider)
}

func getFactorForAmount(amount int) (int, string) {
	lengthOfNumber := int(math.Log10(float64(amount)) + 1)
	switch {
	case lengthOfNumber <= 3:
		return 100, "Hundred"
	case lengthOfNumber <= 6:
		return 1000, "Thousand"
	case lengthOfNumber <= 9:
		return 1000000, "Million"
	case lengthOfNumber <= 12:
		return 1000000000, "Billion"
	default:
		return 1, "Unknown"
	}
}


func calculatePoundUnits(currentString string, amount int) string {
	//If the current string has been set and we have nothing left, we do not want to set zero
	if currentString != "" && amount == 0 {
		return currentString
	}

	return currentString + belowTwenty[amount] + " "
}

func getPoundsString(amount int) string {
	if amount == 1 {
		return "Pound"
	} else {
		return "Pounds"
	}
}
