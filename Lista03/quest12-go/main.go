package main

//ORDENA LISTA

import "fmt"

func ordem_na_parada(x []int) []int {
	var j []int
	ponto := (len(x))

	for i := 0; i < (ponto - 1); i++ {

		for j := 0; j < ((ponto - 1) - i); j++ {

			if x[j] > x[j+1] {

				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	j = x

	return j
}

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var num int

	fmt.Scan(&num)

	seq1 := make([]int, num)

	for i := range seq1 {

		fmt.Scan(&seq1[i])

	}
	fmt.Printf("%v\n", seq1)

	seq1 = ordem_na_parada(seq1)

	fmt.Printf("%v\n", (seq1))
}
