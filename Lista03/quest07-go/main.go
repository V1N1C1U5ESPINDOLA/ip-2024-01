package main

//ACUMULADO DE ELEMENTOS

import "fmt"

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var (
		num, temp, maior, soma1 int
	)
	var (
		seq1 []int
	)

	fmt.Scan(&num)

	for i := 0; i < num; i++ {

		fmt.Scan(&temp)
		seq1 = append(seq1, temp)

	}

	for j := 0; j < num; j++ {

		if seq1[j] > maior {

			maior = seq1[j]

		}
	}

	for h := 0; h <= maior; h++ {
		soma1 = 0
		for j := 0; j < num; j++ {
			if seq1[j] <= h {
				soma1++
			}
		}

		fmt.Printf("(%v) %v\n", h, soma1)
	}

}
