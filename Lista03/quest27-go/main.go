package main

//INTERCALA

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

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var num1, num2 int

	fmt.Scan(&num1, &num2)

	seq1 := make([]int, num1)

	seq2 := make([]int, num2)

	for i := 0; i < num1; i++ {

		fmt.Scan(&seq1[i])
	}
	for j := 0; j < num2; j++ {

		fmt.Scan(&seq2[j])
	}
	var seq3 []int

	seq3 = append(seq1, seq2...)

	seq3 = bubblesort(seq3)

	for l := 0; l < len(seq3); l++ {

		fmt.Printf("\n%v", seq3[l])
	}
}
