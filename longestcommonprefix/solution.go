package main

import (
	"fmt"
)

func longestCommonPrefix(prefixes []string) string {
	var maxl int
	// var maparefijos map[int][]string

	// Obteniendo el max len de todos los str para ir decreciendo e ir analizando
	for _, pr := range prefixes {
		if len(pr) > maxl {
			maxl = len(pr)
		}
	}
	// for _, prf := range prefixes {
	// 	fmt.Println(prf[:len(prf)-1])
	// 	// if
	// 	// res := append(prefijos[i], prf)
	// 	// fmt.Println(res)
	// }
	// fmt.Println(maxl)
	// idea: analizar el largo de todos y del mas largo ir disminuyendo hasta encontrar el prefijo mas largo
	for i := 0; i < maxl; i++ {
		var prefijos []string
		for _, prf := range prefixes {
			// fmt.Println(prf[:len(prf)-i])
			prefijos = append(prefijos, prf[:len(prf)-i])
		}
		fmt.Println(prefijos)
		// maparefijos[i] = prefijos
	}
	return "ok"
}

func main() {
	prefixes := []string{"flower", "flow", "flight"}
	// prefixes := []string{"flower", "flight"}
	retorno := longestCommonPrefix(prefixes)
	fmt.Println(retorno)
}
