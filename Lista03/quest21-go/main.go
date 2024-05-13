package main

import "fmt"

func impoupa(x []int) []int {

	tamanho := len(x)

	for i := 0; i < tamanho; i++ {

		for j := tamanho - 1; j > i; j-- {

			if x[j] < x[j-1] {

				x[j], x[j-1] = x[j-1], x[j]
			}
		}
	}
	var (
		pares, impares []int
	)
	for i := range x {

		if x[i]%2 == 0 {

			pares = append(pares, x[i])
		}
	}
	for i := 0; i < tamanho; i++ {

		if x[tamanho-1-i]%2 != 0 {

			impares = append(impares, x[tamanho-1-i])
		}
	}
	seq3 := append(pares, impares...)

	return seq3
}

func main() {
	fmt.Println("DIGITE OS NÚMEROS")

	var num int

	fmt.Scan(&num)

	seq1 := make([]int, num)

	for i := range seq1 {

		fmt.Scan(&seq1[i])
	}
	seq2 := impoupa(seq1)

	for j := range seq2 {
		fmt.Printf("\n%v ", seq2[j])
	}
}
