package main

// CÁLCULO DA RAIZ QUADRADA

import (
	"fmt"
)

func main() {
	fmt.Println("DIGITE OS NÚMEROS:")

	var num, margin float64
	fmt.Scan(&num, &margin)

	xn := 1.0

	for {
		x := 0.5 * (xn + (num / xn))

		err := (x*x - num)

		fmt.Printf("r: %.9f erro: %.9f\n", x, err)

		xn = x
		if err < margin {
			break
		}
	}
}
