package main

//Exercício 002

import "fmt"

func main() {

	fmt.Println("DIGITE CINCO VALORES A SEREM SOMADOS")
	const num = 5
	var values [num]int
	var soma int
	for i := range values {
		fmt.Scan(&values[i])
		soma += values[i]
		fmt.Printf("+ %d ", values[i])

	}
	fmt.Printf("= %d", soma)

}
