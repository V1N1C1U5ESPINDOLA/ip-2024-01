package main

//SÉRIE DE PARES

import "fmt"

func main() {

	fmt.Println("DIGITE UM NÚMERO PAR E O NÚMERO DE REPETIÇÕES")
	var (
		numpar, numrep, i int32
	)
	fmt.Scan(&numpar, &numrep)

	if numpar%2.0 != 0 {
		fmt.Println("O PRIMEIRO NÚMERO NÃO É PAR")

	} else if numpar%2.0 == 0 {
		for i = 0; i < numrep; i++ {
			fmt.Println(numpar + 2*i)
		}

	}
}
