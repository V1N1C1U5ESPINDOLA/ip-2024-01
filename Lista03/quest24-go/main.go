package main

import "fmt"

func countingSort(x []int) []int {

	n := len(x)
	m := 0

	for i := range x {

		if x[i] > m {

			m = x[i]
		}
	}
	vcount := make([]int, m+1)
	vord := make([]int, n)

	for i := range x {

		vcount[x[i]]++
	}
	total := 0

	for i := range vcount {

		countanterior := vcount[i]

		vcount[i] = total

		total += countanterior
	}
	for i := range x {

		vord[vcount[x[i]]] = x[i]

		vcount[x[i]]++
	}
	return vord
}

func main() {

	var seq2 [][]int

	fmt.Println("DIGITE OS NÚMEROS")

	for {
		var num int

		fmt.Scan(&num)

		if num == 0 {
			break
		}
		seq1 := make([]int, num)

		for i := range seq1 {

			fmt.Scan(&seq1[i])
		}
		seq2 = append(seq2, countingSort(seq1))
	}
	for i := range seq2 {

		fmt.Println(seq2[i])
	}
}
