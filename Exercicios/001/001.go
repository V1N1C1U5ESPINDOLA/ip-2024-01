package main

//Exercício 01

import "fmt"

func main() {

	fmt.Println("DIGITE SUAS TRÊS NOTAS")
	const num = 3

	var (
		notas [num]float32
	)
	var (
		soma float32
	)
	soma = 0.0
	for i := 0; i < num; i++ {
		fmt.Scan(&notas[i])
		soma += notas[i]

	}
	media := soma / num
	fmt.Printf("SUAS NOTA FORAM RESPECTIVAMENTE:\n")
	for j, v := range notas {
		fmt.Printf("NOTA %d = %.2f\n", (j + 1), v)
	}
	fmt.Printf("SUA MÉDIA É = %.2f", media)
}
