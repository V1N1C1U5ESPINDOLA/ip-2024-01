package main

//REAJUSTE DE SALÁRIO

import "fmt"

func main() {

	fmt.Println("DIGITE SEU SALÁRIO")
	var salario float32
	fmt.Scan(&salario)

	if salario <= 300 {
		salario += (salario / 2)
	} else if salario > 300 {
		salario += (salario * 0.3)
	}
	fmt.Println("SALÁRIO COM REAJUSTE =", salario)
}
