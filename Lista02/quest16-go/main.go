package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Digite um número inteiro maior que zero:")
	fmt.Scanln(&num)

	for hip := 1; hip <= num; hip++ {
		for cat1 := 1; cat1 < hip; cat1++ {
			for cat2 := cat1; cat2 < hip; cat2++ {
				if cat1*cat1+cat2*cat2 == hip*hip {
					fmt.Printf("HIPOTENUSA = %d, CATETOS %d E %d\n", hip, cat1, cat2)
				}
			}
		}
	}
}
