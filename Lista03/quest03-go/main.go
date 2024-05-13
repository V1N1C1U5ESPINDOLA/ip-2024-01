package main

//IMPRIMIR UM VETOR NA ORDEM INVERSA

import "fmt"

func trocatroca(x []int) []int {

	var y []int
	for i := 0; i < len(x); i++ {
		y = append(y, x[len(x)-1-i])
	}
	return y
}

func main() {

	var (
		num, temp int
	)

	var (
		seq1 []int
	)

	fmt.Println("DIGITE OS NÚMEROS")

	fmt.Scan(&num)

	for j := 0; j < num; j++ {

		fmt.Scan(&temp)

		seq1 = append(seq1, temp)
	}

	seq2 := trocatroca(seq1)

	fmt.Printf("%v", seq2)

}
