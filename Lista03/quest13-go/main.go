package main

//MAIOR FREQUÊNCIA

import "fmt"

func frequentemente(x []int) (int, int) {

	contador := 0

	maisfreak := 0

	for i := 0; i < len(x); i++ {

		frequente := 1

		for j := i + 1; j < len(x); j++ {

			if x[i] == x[j] {

				frequente++
			}
		}
		if frequente > contador {

			contador = frequente

			maisfreak = x[i]
		}
	}
	return maisfreak, contador
}

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var (
		num int
	)

	fmt.Scan(&num)

	seq1 := make([]int, num)

	for i := range seq1 {

		fmt.Scan(&seq1[i])
	}
	a, b := frequentemente(seq1)

	fmt.Printf("%v\n%v", a, b)
}
