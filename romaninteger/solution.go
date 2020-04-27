package romaininteger

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// RomanToInt Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
func RomanToInt(roman string) int {
	roman = strings.ToUpper(roman)
	// Analisis de las excepciones y se reemplaza por el valor con un *
	if strings.Contains(roman, "CM") {
		roman = strings.ReplaceAll(roman, "CM", "900*")
	}
	if strings.Contains(roman, "XC") {
		roman = strings.ReplaceAll(roman, "XC", "90*")
	}
	if strings.Contains(roman, "IX") {
		roman = strings.ReplaceAll(roman, "IX", "9*")
	}
	if strings.Contains(roman, "IV") {
		roman = strings.ReplaceAll(roman, "IV", "4*")
	}
	if strings.Contains(roman, "XL") {
		roman = strings.ReplaceAll(roman, "XL", "40*")
	}
	if strings.Contains(roman, "CD") {
		roman = strings.ReplaceAll(roman, "CD", "400*")
	}

	// Recorrer el resto del string y analizando los casos, reemplazando el caracter por el valor con *
	runes := []rune(roman)
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			if strings.Contains(roman, "I") {
				roman = strings.ReplaceAll(roman, "I", "1*")
			}
			if strings.Contains(roman, "V") {
				roman = strings.ReplaceAll(roman, "V", "5*")
			}
			if strings.Contains(roman, "X") {
				roman = strings.ReplaceAll(roman, "X", "10*")
			}
			if strings.Contains(roman, "L") {
				roman = strings.ReplaceAll(roman, "L", "50*")
			}
			if strings.Contains(roman, "C") {
				roman = strings.ReplaceAll(roman, "C", "100*")
			}
			if strings.Contains(roman, "D") {
				roman = strings.ReplaceAll(roman, "D", "500*")
			}
			if strings.Contains(roman, "M") {
				roman = strings.ReplaceAll(roman, "M", "1000*")
			}
		}
		// fmt.Println(string(runes[i]))
	}
	// Se elimina el ultimo caracter *
	r := []rune(roman)
	roman = string(r[:len(r)-1])

	// Se realiza split para convertir el string en slice de strings y
	contenidoFinal := strings.Split(roman, "*")
	var total int
	fmt.Println(contenidoFinal)
	for _, numero := range contenidoFinal {
		// por cada elemento convertir a entero
		parcial, err := strconv.Atoi(numero)
		if err != nil {
			fmt.Println("Ocurrió un error")
		}
		// Se suman los enteros para obtener el valor final
		total += parcial
	}
	return total
}
