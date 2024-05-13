package main

//CPF

import "fmt"

func cpfcancelado(x []int) string {

	var somab1, somab2 int

	for i := 0; i < 9; i++ {

		somab2 = ((x[(9 - i - 1)]) * (i + 1)) + somab2

		somab1 = (x[i] * (i + 1)) + somab1
	}
	b1 := somab1 % 11

	b2 := somab2 % 11

	if b2 == 10 {

		b2 = 0
	}
	if b1 == 10 {

		b1 = 0
	}
	if b1 == x[9] && b2 == x[10] {

		return "VÁLIDO"

	} else {

		return "INVÁLIDO"
	}
}
func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var testes int

	var seq2 []string

	fmt.Scan(&testes)

	for i := 0; i < testes; i++ {

		seq1 := make([]int, 11)

		for j := range seq1 {

			fmt.Scan(&seq1[j])
		}
		s := cpfcancelado(seq1)

		seq2 = append(seq2, s)
	}
	fmt.Printf("%s", seq2)

	for i := 0; i < testes; i++ {

		fmt.Printf("\n CPF %v", seq2[i])
	}
}
