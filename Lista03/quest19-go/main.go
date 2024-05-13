package main

//ELEMENTOS UNICOS

import "fmt"

func bubblesort(l []int) []int {

	n := len(l)

	for i := 0; i < n-1; i++ {

		for j := n - 1; j > i; j-- {

			if l[j-1] > l[j] {

				l[j-1], l[j] = l[j], l[j-1]
			}
		}
	}
	return l
}

func semrep(x []int) []int {

	var seq3 []int

	freq := make(map[int]int)

	for _, i := range x {

		freq[i]++
	}
	for j := range freq {

		seq3 = append(seq3, j)
	}

	return seq3
}

func main() {

	var seq2 []int

	fmt.Println("DIGITE OS NÚMEROS")

	var num int

	fmt.Scan(&num)

	seq1 := make([]int, num)

	for i := range seq1 {

		fmt.Scan(&seq1[i])
	}
	fmt.Printf("\n")

	seq2 = bubblesort(semrep(seq1))

	for j := range seq2 {

		fmt.Printf("\n%v", seq2[j])
	}
}
