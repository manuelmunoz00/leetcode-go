package romaininteger

var mapsroman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanToInt Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
func RomanToInt(roman string) int {
	var total int
	var next string
	for i := 0; i < len(roman); i++ {
		var noob bool = false
		x := i + 1
		c := string(roman[i])
		if x == len(roman) {
			x--
		}
		next = string(roman[x])
		if c == "I" && (next == "V" || next == "X") {
			total = total - 1
			noob = true
		}
		if c == "X" && (next == "L" || next == "C") {
			total = total - 10
			noob = true
		}
		if c == "C" && (next == "D" || next == "M") {
			total = total - 100
			noob = true
		}
		if !noob {
			total += mapsroman[c]
		}
	}
	return total
}
