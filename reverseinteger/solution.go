package reverseinteger

import (
	"fmt"
	"math"
	"strconv"
)

// Reverse func
func Reverse(numero int) int {
	strNumero := strconv.Itoa(numero)
	var sliceDigitos []int
	var resultado int
	var reverseStr string
	var negativo bool
	if numero < 0 {
		numero = numero * -1
		strNumero = strconv.Itoa(numero)
		negativo = true
	}
	for i := 0; i < len(strNumero); i++ {
		if i == 0 {
			resultado = numero % 10
			if resultado == 0 {
				continue
			}
			sliceDigitos = append(sliceDigitos, resultado)
			reverseStr = reverseStr + strconv.Itoa(resultado)
		} else {
			numero = numero / 10
			resultado = numero % 10
			sliceDigitos = append(sliceDigitos, resultado)
			reverseStr = reverseStr + strconv.Itoa(resultado)
		}
	}
	numOk, err := strconv.Atoi(reverseStr)
	if err != nil {
		fmt.Println("error!")
	}

	if numOk > math.MaxInt32 {
		return 0
	} else if numOk < math.MinInt32 {
		return 0
	}

	if negativo {
		return numOk * -1
	}

	return numOk
}

// func main() {
// 	numero := 120
// 	fmt.Println(Reverse(numero))
// }
