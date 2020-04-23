package reverseinteger

import (
	"strconv"
	"testing"
)

var isReverseTests = []struct {
	numero, esperado int
}{
	{123456, 654321},
	{-123, -321},
	{654321, 123456},
	{1534236469, 0},
	{-2147483648, 0},
}

func TestReverse(t *testing.T) {
	for _, tt := range isReverseTests {
		t.Run(strconv.Itoa(tt.numero), func(t *testing.T) {
			got := Reverse(tt.numero)
			if got != tt.esperado {
				t.Errorf("Expected: %v, got: %v", tt.esperado, got)
			}
		})
	}
}
