package main

//SOMA ELEMENTOS DO VETOR

import "fmt"

func main() {

	var (
		num, temp, soma int
	)
	var seq1 []int

	fmt.Println("DIGITE OS NÚMEROS")

	fmt.Scan(&num)

	for i := 0; i < num; i++ {

		fmt.Scan(&temp)

		seq1 = append(seq1, temp)

	}
	for j := 0; j < len(seq1); j++ {

		soma = soma + seq1[j]

	}
	fmt.Printf("%v", soma)
}
