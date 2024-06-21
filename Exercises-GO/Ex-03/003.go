package main

// Exercício 003

import "fmt"

func main() {

	fmt.Println("DIGITE DEZ NÚMEROS REAIS")
	const N = 10
	var num [N]float32
	for i := 0; i < N; i++ {

		fmt.Scan(&num[i])

	}

	for j := 0; j < N; j++ {

		new := N - j
		fmt.Printf("%.0f ", num[new-1])
	}
}
