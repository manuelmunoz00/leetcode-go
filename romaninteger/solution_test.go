package romaininteger

import (
	"strconv"
	"testing"
)

var isIntegerRoman = []struct {
	roman  string
	numero int
}{
	{"IV", 4},
	{"III", 3},
	{"MM", 2000},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"MCCXLIX", 1249},
	{"CMXCIX", 999},
}

func TestRomanToInt(t *testing.T) {
	for _, tt := range isIntegerRoman {
		t.Run(strconv.Itoa(tt.numero), func(t *testing.T) {
			got := RomanToInt(tt.roman)
			if got != tt.numero {
				t.Errorf("Expected: %v, got: %v", tt.numero, got)
			}
		})
	}
}
