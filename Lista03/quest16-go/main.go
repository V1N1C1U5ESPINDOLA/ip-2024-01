package main

//AULA CANCELADA

import "fmt"

func greve(y int, x []int) (string, []int) {

	var seq1 []int
	soma := 0

	for i := (len(x) - 1); i >= 0; i = i - 1 {

		if x[i] <= 0 {

			soma += 1

			seq1 = append(seq1, (i + 1))
		}
	}
	if soma < y {

		return "SIM", seq1

	} else {

		return "NÃO", seq1
	}
}

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var num, minimo int

	fmt.Scan(&num, &minimo)

	seq2 := make([]int, num)

	for i := range seq2 {

		fmt.Scan(&seq2[i])
	}
	a, b := greve(minimo, seq2)

	fmt.Printf("\n%v", a)

	if a == "NÃO" {

		for i := 0; i < len(b); i++ {

			fmt.Printf("\n%v", b[i])
		}
	}
}
