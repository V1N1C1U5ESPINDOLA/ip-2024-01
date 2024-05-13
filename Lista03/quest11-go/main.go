package main

//INVERTE VETOR

import "fmt"

func mundo_invertido(x []int) []int {

	var seq2 []int

	for i := 0; i < len(x); i++ {

		seq2 = append(seq2, x[(len(x)-1)-i])

	}
	return seq2
}

func extremos(x []int) (int, int) {

	maiornum, menornum := x[0], x[0]

	for i := 0; i < len(x); i++ {

		if x[i] > maiornum {

			maiornum = x[i]

		}

		if x[i] < menornum {

			menornum = x[i]

		}
	}
	return maiornum, menornum
}

func main() {
	var (
		num int
	)

	fmt.Println("DIGITE OS NÚMEROS")

	fmt.Scan(&num)

	seq1 := make([]int, num)

	for i := range seq1 {

		fmt.Scan(&seq1[i])

	}
	a, b := extremos(seq1)

	fmt.Printf("%v\n%v\n%v\n%v", seq1, mundo_invertido(seq1), a, b)

}
