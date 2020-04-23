package main

import (
	"fmt"
)

// Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
func romanToInt(roman string) int {
	// var countii, countv, countx, countl, countc, countd, countm int
	// var countiv, countix, countxl, countxc, countcd, countcm int
	runes := []rune(roman)
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}
	return 99999
}

func main() {
	fmt.Println(romanToInt("MMC"))
}
