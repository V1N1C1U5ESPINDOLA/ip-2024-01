package main

//NÚMERO PRIMO

import (
	"fmt"
)

func main() {

	fmt.Println("DIGITE UM NÚMERO INTEIRO")

	var (
		num, condição int
	)
	fmt.Scan(&num)
	if num <= 0 {
		fmt.Println("NÚMERO INVÁLIDO")
		return

	}
	if num == 2 {

		fmt.Printf("O NÚMERO %d É PRIMO", num)

		return
	}

	for i := 2; i <= num/2; i++ {

		if num%i == 0 {

			condição = 0
			break
		} else {

			condição = 1
		}

	}
	if condição == 0 {

		fmt.Printf("O NÚMERO %d NÃO É PRIMO", num)

	} else if condição == 1 {

		fmt.Printf("O NÚMERO %d É PRIMO", num)

	}
}
