package main

//SOMA DE PROGRESSÃO ARITIMÉTICA

import "fmt"

func main() {

	fmt.Println("DIGITE OS TRÊS ALGARISMOS")

	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)
	var (
		PA, soma int
	)

	for i := 0; i < c; i++ {
		PA = a + b*i
		soma = PA + soma
	}
	fmt.Printf("VALOR DA PA %v", soma)
}
