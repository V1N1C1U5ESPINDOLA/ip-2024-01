package main

//SOMATÓRIO SIMPLES

import "fmt"

func main() {

	fmt.Println("DIGITE UM NÚMERO INTEIRO E POSITIVO")
	var num float32
	fmt.Scan(&num)
	if num < 1 {
		fmt.Println("NÚMERO INVÁLIDO")
	} else if num >= 1 {
		var (
			k, s float32
		)

		for k = 1; k <= num; k++ {

			s = (1 / k) + s

		}
		fmt.Println(s)
	}

}
