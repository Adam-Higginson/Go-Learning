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
	return fmt.Sprintf("%s%s And %s",
		convertToChequeFormatRecurse("", poundsAmount),
		getPoundsString(poundsAmount),
		convertPenceAmount(penceAmount)), nil
}

func convertToChequeFormatRecurse(currentString string, poundsAmount int) string {
	fmt.Printf("CurrentString=%s PoundsAmount=%d\n", currentString, poundsAmount)
	if poundsAmount < 20 {
		return calculatePoundUnits(currentString, poundsAmount)
	} else if poundsAmount < 100 {
		return calculatePoundTens(currentString, poundsAmount)
	} else if poundsAmount < 1000 {
		return calculatePoundHundreds(currentString, poundsAmount)
	} else {
		return calculatePoundThousandsAndGreater(currentString, poundsAmount)
	}
}

func calculatePoundUnits(currentString string, poundsAmount int) string {
	//If the current string has been set and we have nothing left, we do not want to set zero
	if currentString != "" && poundsAmount == 0 {
		return currentString
	}

	return currentString + belowTwenty[poundsAmount] + " "
}

func calculatePoundTens(currentString string, poundsAmount int) string {
	currentDigit := poundsAmount / 10
	convertedToString := currentString + belowHundred[currentDigit] + " "
	return convertToChequeFormatRecurse(convertedToString, poundsAmount % 10)
}

func calculatePoundHundreds(currentString string, poundsAmount int) string {
	currentDigit := poundsAmount / 100
	convertedToString := currentString + belowTwenty[currentDigit] + " Hundred "

	return convertToChequeFormatRecurse(convertedToString, poundsAmount % 100)
}

func calculatePoundThousandsAndGreater(currentString string, poundsAmount int) string {
	divider, name := getFactorForPoundsAmount(poundsAmount)

	currentDigit := poundsAmount / divider
	prefix := convertToChequeFormatRecurse("", currentDigit)
	convertedToString := currentString + prefix + name + " "

	return convertToChequeFormatRecurse(convertedToString, poundsAmount % divider)
}


func getFactorForPoundsAmount(poundsAmount int) (int, string) {
	lengthOfNumber := int(math.Log10(float64(poundsAmount)) + 1)
	if lengthOfNumber <= 6 {
		return 1000, "Thousand"
	} else if lengthOfNumber <= 9 {
		return 1000000, "Million"
	} else if lengthOfNumber <= 12 {
		return 1000000000, "Billion"
	}

	return 1, "Unknown"
}

func convertPenceAmount(penceAmount int) string {
	if penceAmount < 20 {
		return belowTwenty[penceAmount] + " Pence"
	} else {
		return belowHundred[(penceAmount / 10)] + " " + belowTwenty[penceAmount % 10] + " Pence"
	}
}

func getPoundsString(amount int) string {
	if amount == 1 {
		return "Pound"
	} else {
		return "Pounds"
	}
}
