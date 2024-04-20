package main

//GERADOR DE TABUADA

import (
	"fmt"
)

func main() {

	fmt.Println("DIGITE NÚMEROS DE 0 À 9")
	var (
		num1, num2, num3, num4, i, e, u, o, r1, r2, r3, r4 float64
	)

	fmt.Scan(&num1, &num2, &num3, &num4)
	fmt.Print("\n")
	fmt.Printf("TABUADA DE SOMA:\n")
	fmt.Print("\n")
	for i = 0; i <= num3; i++ {
		r1 = num1 + num2 + (num4 * i)
		fmt.Printf("%.2f+%.2f=%.2f\n", num1, num2+(num4*i), r1)
	}
	fmt.Print("\n")
	fmt.Printf("TABUADA DE SUBTRAÇÃO:\n")
	fmt.Print("\n")
	for e = 0; e <= num3; e++ {
		r2 = num1 - num2 + (num4 * e)
		fmt.Printf("%.2f-%.2f=%.2f\n", num1, num2+(num4*e), r2)
	}
	fmt.Print("\n")
	fmt.Printf("TABUADA DE MULTIPLICAÇÃO:\n")
	fmt.Print("\n")
	for u = 0; u <= num3; u++ {
		r3 = num1 * (num2 + (num4 * u))
		fmt.Printf("%.2f*%.2f=%.2f\n", num1, (num2 + (num4 * u)), r3)
	}
	fmt.Print("\n")
	fmt.Printf("TABUADA DE DIVISÃO:\n")
	fmt.Print("\n")
	for o = 0; o <= num3; o++ {
		r4 = num1 / (num2 + (num4 * o))
		fmt.Printf("%.2f/%.2f=%.2f\n", num1, (num2 + (num4 * o)), r4)

	}
	fmt.Print("\n")
}
