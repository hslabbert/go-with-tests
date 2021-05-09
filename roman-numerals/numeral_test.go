package numerals

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 54, Roman: "LIV"},
		{Arabic: 66, Roman: "LXVI"},
		{Arabic: 89, Roman: "LXXXIX"},
		{Arabic: 97, Roman: "XCVII"},
		{Arabic: 201, Roman: "CCI"},
		{Arabic: 401, Roman: "CDI"},
		{Arabic: 601, Roman: "DCI"},
		{Arabic: 801, Roman: "DCCCI"},
		{Arabic: 901, Roman: "CMI"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			assertNumerals(t, got, want)
		})
	}
}

func assertNumerals(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
