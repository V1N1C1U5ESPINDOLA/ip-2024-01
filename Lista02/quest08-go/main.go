package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("DIGITE A QUANTIDADE DE TIMES")
	fmt.Scan(&num)

	if num < 2 {
		fmt.Println("CAMPEONATO INVÁLIDO!")
		return
	}

	fmt.Println("FINAIS POSSÍVEIS:")
	numdejogos := 1
	for i := 1; i <= num; i++ {
		for j := i + 1; j <= num; j++ {
			fmt.Printf("Final %d: Time%d X Time%d\n", numdejogos, i, j)
			numdejogos = numdejogos + 1
		}
	}
}
