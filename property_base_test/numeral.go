package propertybasetest

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

var allRomanMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func ConvertToArabic(roman string) int {
	arabicNum := 0
	romanStr := roman

	for len(romanStr) > 0 {
		if len(romanStr) >= 2 {
			if val, ok := allRomanMap[romanStr[0:2]]; ok {
				arabicNum += val
				romanStr = romanStr[2:]
			} else {
				arabicNum += allRomanMap[romanStr[0:1]]
				romanStr = romanStr[1:]
			}
		} else {
			arabicNum += allRomanMap[romanStr]
			romanStr = romanStr[1:]
		}
	}

	return arabicNum
}
