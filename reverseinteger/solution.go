package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func reverse(number int) int {

	resultado := number % 10
	//fmt.Println(resultado)
	return resultado
}

func main() {
	// numero := reverse(3452)
	// fmt.Println(numero)

	numero2 := 102500
	strNumero := strconv.Itoa(numero2)
	fmt.Println(strNumero)
	fmt.Println(reflect.TypeOf(strNumero))
	intNumero, error := strconv.Atoi(strNumero)
	if error != nil {
		fmt.Println("Se produjo un error")
	}
	fmt.Println(intNumero)
	fmt.Println(reflect.TypeOf(intNumero))
	// var elementos []string
	// var n int(rango de -2147483647 a 2147483647)
	fmt.Println(len(strNumero))
	for _, numero3 := range strNumero {
		fmt.Println(string(numero3))
		// fmt.Println(reflect.TypeOf(int(numero)))
		// append(elementos, string(numero))
	}
	// fmt.Println(elementos)
}
