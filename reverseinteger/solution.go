package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(numero int) int {
	if numero > math.MaxInt32 {
		return 0
	} else if numero < math.MinInt32 {
		return 0
	}

	strNumero := strconv.Itoa(numero)
	var sliceDigitos []int
	var resultado int
	var reverseStr string
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
	return numOk
}

func main() {
	numero := 2147483647
	fmt.Println(reverse(numero))
}
