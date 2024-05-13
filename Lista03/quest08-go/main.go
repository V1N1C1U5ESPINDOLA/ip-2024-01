package main

//CONVERSÃO DE DECIMAL PARA BINÁRIO

import "fmt"

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var (
		num, temp int
	)
	var (
		seq1 []int
	)

	fmt.Scan(&num)

	for i := 0; i < num; i++ {

		fmt.Scan(&temp)

		seq1 = append(seq1, temp)

	}

	for j := 0; j < len(seq1); j++ {

		fmt.Printf("%b\n", seq1[j])
	}
}
