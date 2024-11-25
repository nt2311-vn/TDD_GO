package propertybasetest

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"50 gets converted to L", 50, "L"},
		{"100 gets converted to C", 100, "C"},
		{"90 gets converted to XC", 90, "XC"},
		{"400 gets converted to CD", 400, "CD"},
		{"500 gets converted to D", 500, "D"},
		{"900 gets converted to CM", 900, "CM"},
		{"1000 gets converted to M", 1000, "M"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"2014 gets converted to MMXIV", 2014, "MMXIV"},
		{"1006 gets converted to MVI", 1006, "MVI"},
		{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	cases := []struct {
		Description string
		Roman       string
		Value       int
	}{
		{"1 gets converted to I", "I", 1},
		{"2 gets converted to II", "II", 2},
		{"3 gets converted to III", "III", 3},
		{"4 gets converted to IV (can't repeat more than 3 times)", "IV", 4},
		{"5 gets converted to V", "V", 5},
		{"9 gets converted to IX", "IX", 9},
		{"10 gets converted to X", "X", 10},
		{"14 gets converted to XIV", "XIV", 14},
		{"18 gets converted to XVIII", "XVIII", 18},
		{"20 gets converted to XX", "XX", 20},
		{"39 gets converted to XXXIX", "XXXIX", 39},
		{"40 gets converted to XL", "XL", 40},
		{"47 gets converted to XLVII", "XLVII", 47},
		{"50 gets converted to L", "L", 50},
		{"100 gets converted to C", "C", 100},
		{"90 gets converted to XC", "XC", 90},
		{"400 gets converted to CD", "CD", 400},
		{"500 gets converted to D", "D", 500},
		{"900 gets converted to CM", "CM", 900},
		{"1000 gets converted to M", "M", 1000},
		{"1984 gets converted to MCMLXXXIV", "MCMLXXXIV", 1984},
		{"3999 gets converted to MMMCMXCIX", "MMMCMXCIX", 3999},
		{"2014 gets converted to MMXIV", "MMXIV", 2014},
		{"1006 gets converted to MVI", "MVI", 1006},
		{"798 gets converted to DCCXCVIII", "DCCXCVIII", 798},
	}
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Value), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Value {
				t.Errorf("got %d, want %d", got, test.Value)
			}
		})
	}
}
