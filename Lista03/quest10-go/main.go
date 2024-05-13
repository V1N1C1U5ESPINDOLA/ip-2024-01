package main

//FREQUENCIA MAIOR

import (
	"fmt"
)

func frequencia(x []int) (int, int) {

	var (
		ultimo, vezes int
	)

	for i := 0; i < len(x); i++ {

		if x[len(x)-1] == x[i] {

			vezes++

		}
		ultimo = x[len(x)-1]
	}
	return ultimo, vezes
}

func maiornota(x []int) (int, int) {

	var (
		maior, indice int
	)

	for i := 0; i < len(x); i++ {

		if maior < x[i] {

			maior = x[i]

			indice = i

		}

	}
	return maior, indice
}

func main() {

	var num int

	fmt.Println("DIGITE OS NÚMEROS")

	fmt.Scan(&num)

	seq1 := make([]int, num)

	for i := range seq1 {

		fmt.Scan(&seq1[i])

	}

	a, b := frequencia(seq1)

	c, d := maiornota(seq1)

	fmt.Printf("NOTA %v, %v VEZES\nNOTA %v, INDICE %v", a, b, c, d)
}
