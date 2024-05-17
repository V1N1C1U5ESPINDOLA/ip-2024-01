package main

//UNIÃO E INTERSECÇÃO DE CONJUNTOS

import "fmt"

func inter(x, y []int) []int {
	var seccao []int

	mapa := make(map[int]bool)
	for _, n := range x {
		mapa[n] = true
	}
	for _, n := range y {
		if _, sim := mapa[n]; sim {
			seccao = append(seccao, n)
		}
	}
	return seccao
}
func org(x []int) []int {

	cont := 0
	var slice1 []int

	for i := 0; i < len(x); i++ {
		cont = 0
		for j := 0; j < len(slice1); j++ {
			if len(slice1) == 0 {
				slice1 = append(slice1, x[i])
			} else if x[i] == slice1[j] {
				cont++
			}
		}
		if cont == 0 {
			slice1 = append(slice1, x[i])
		}
	}
	return slice1
}

func main() {

	fmt.Println("DIGITE OS NÚMEROS")
	var N1, N2 int
	for {
		var num1 int
		fmt.Scan(&num1)
		if num1 > 0 && num1 < 1000 {
			N1 = num1
			break
		}
	}
	for {
		var num2 int
		fmt.Scan(&num2)
		if num2 > 0 && num2 < 1000 {
			N2 = num2
			break
		}
	}
	seq1 := make([]int, N1)
	seq2 := make([]int, N2)
	for i := 0; i < N1; i++ {

		fmt.Scan(&seq1[i])
	}
	for j := 0; j < N2; j++ {

		fmt.Scan(&seq2[j])
	}
	var seq3 []int
	seq1 = org(seq1)
	seq2 = org(seq2)
	seq3 = append(seq2, seq1...)
	seq3 = (org(seq3))
	fmt.Printf("%v\n", seq3)
	fmt.Printf("%v", inter(seq1, seq2))

}
