package main

//MÍNIMO MULTIPLO COMUM

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func mmc(a, b int) int {
	return a * b / mdc(a, b)
}
func mmc3(a, b, c int) int {
	return mmc(a, mmc(b, c))
}

func main() {
	fmt.Println("DIGITE TRÊS NÚMEROS INTEIROS")
	var (
		num1, num2, num3 int
	)
	fmt.Scan(&num1, &num2, &num3)

}
