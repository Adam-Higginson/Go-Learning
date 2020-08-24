package cheque

import (
	"testing"
)

var expectedBelowTwenty = [20]string {
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


func TestConvertWithOnePound(t *testing.T) {
	converted, err := ConvertToChequeFormat(100)
	var expected = "One Pound And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestConvertWithTwoPounds(t *testing.T) {
	converted, err := ConvertToChequeFormat(200)
	var expected = "Two Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestConvertWithThreePounds(t *testing.T) {
	converted, err := ConvertToChequeFormat(300)
	var expected = "Three Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestConvertBelowTwentyPounds(t *testing.T) {
	for i := 0; i < 20; i++ {
		converted, err := ConvertToChequeFormat(i * 100)
		assertNilError(t, err)
		if i == 1 {
			assertEquals(t, expectedBelowTwenty[i] + " Pound And Zero Pence", converted)
		} else {
			assertEquals(t, expectedBelowTwenty[i] + " Pounds And Zero Pence", converted)
		}
	}
}

func TestHandlePoundsAndPence(t *testing.T) {
	converted, err := ConvertToChequeFormat(315)
	var expected = "Three Pounds And Fifteen Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleGreaterThanTwenty(t *testing.T) {
	converted, err := ConvertToChequeFormat(2100)
	var expected = "Twenty One Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleGreaterThanTwentyWithPence(t *testing.T) {
	converted, err := ConvertToChequeFormat(2105)
	var expected = "Twenty One Pounds And Five Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleWithPenceAndNoZero(t *testing.T) {
	converted, err := ConvertToChequeFormat(630)
	var expected = "Six Pounds And Thirty Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePenceAmountOnly(t *testing.T) {
	converted, err := ConvertToChequeFormat(17)
	var expected = "Zero Pounds And Seventeen Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePenceGreaterThanTwenty(t *testing.T) {
	converted, err := ConvertToChequeFormat(99)
	var expected = "Zero Pounds And Ninety Nine Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePoundsGreaterThanTwentyPenceGreaterThanTwenty(t *testing.T) {
	converted, err := ConvertToChequeFormat(9999)
	var expected = "Ninety Nine Pounds And Ninety Nine Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePoundsGreaterThanHundred(t *testing.T) {
	converted, err := ConvertToChequeFormat(15000)
	var expected = "One Hundred Fifty Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePoundsExactlyHundred(t *testing.T) {
	converted, err := ConvertToChequeFormat(10000)
	var expected = "One Hundred Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePoundsJustLessThanThousand(t *testing.T) {
	converted, err := ConvertToChequeFormat(99999)
	var expected = "Nine Hundred Ninety Nine Pounds And Ninety Nine Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePoundsOfOneThousand(t *testing.T) {
	converted, err := ConvertToChequeFormat(100000)
	var expected = "One Thousand Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandlePoundsOfThousandNotExact(t *testing.T) {
	converted, err := ConvertToChequeFormat(115000)
	var expected = "One Thousand One Hundred Fifty Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleTenThousand(t *testing.T) {
	converted, err := ConvertToChequeFormat(1000000)
	var expected = "Ten Thousand Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleTensOfThousandsAndPence(t *testing.T) {
	converted, err := ConvertToChequeFormat(1200052)
	var expected = "Twelve Thousand Pounds And Fifty Two Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleThousandsGreaterThanTwenty(t *testing.T) {
	converted, err := ConvertToChequeFormat(2100000)
	var expected = "Twenty One Thousand Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleThousandsGreaterThanHundred(t *testing.T) {
	converted, err := ConvertToChequeFormat(12100000)
	var expected = "One Hundred Twenty One Thousand Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleThousandsJustLessThanAMillion(t *testing.T) {
	converted, err := ConvertToChequeFormat(99999999)
	var expected = "Nine Hundred Ninety Nine Thousand Nine Hundred Ninety Nine Pounds And Ninety Nine Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleOneMillion(t *testing.T) {
	converted, err := ConvertToChequeFormat(100000000)
	var expected = "One Million Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleFiveMillion(t *testing.T) {
	converted, err := ConvertToChequeFormat(500000000)
	var expected = "Five Million Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}


func TestHandleComplexMillion(t *testing.T) {
	converted, err := ConvertToChequeFormat(543200000)
	var expected = "Five Million Four Hundred Thirty Two Thousand Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleSevenBillion(t *testing.T) {
	converted, err := ConvertToChequeFormat(700000000000)
	var expected = "Seven Billion Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleComplexBillion(t *testing.T) {
	converted, err := ConvertToChequeFormat(765432102345)
	var expected = "Seven Billion Six Hundred Fifty Four Million Three Hundred Twenty One Thousand Twenty Three Pounds And Forty Five Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleThousandWithNoHundred(t *testing.T) {
	converted, err := ConvertToChequeFormat(1202352)
	var expected = "Twelve Thousand Twenty Three Pounds And Fifty Two Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleNegativeNumbers(t *testing.T) {
	converted, err := ConvertToChequeFormat(-1)
	if err == nil {
		t.Errorf("Error was nil but expected error for negative amount!")
	} else {
		assertEquals(t, "", converted)
	}
}

func TestExampleGivenInSlack(t *testing.T) {
	converted, err := ConvertToChequeFormat(1021636011)
	var expected = "Ten Million Two Hundred Sixteen Thousand Three Hundred Sixty Pounds And Eleven Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)
}

func TestHandleZeroCase(t *testing.T) {
	converted, err := ConvertToChequeFormat(0)
	var expected = "Zero Pounds And Zero Pence"
	assertNilError(t, err)
	assertEquals(t, expected, converted)

}

func assertNilError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error was not nil!")
	}
}

func assertEquals(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Converted was incorrect, got: %s, wanted: %s", actual, expected)
	}
}