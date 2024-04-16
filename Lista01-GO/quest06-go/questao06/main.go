package main

// CONVERÇÃO DE TEMPERATURA

import "fmt"

func conversao(f float32) float32 {
	return (5 * (f - 32)) / 9
}

func main() {

	fmt.Println("DIGITE O NÚMERO DE TEMPERATURAS A SEREM CONVERTIDAS")
	var numtemp int
	fmt.Scan(&numtemp)
	fmt.Scan()

	var i int
	fmt.Println("DIGITE AS TEMPERATURAS EM FAHRENHEIT")
	for i = 0; i < numtemp; i++ {
		var tempf float32
		fmt.Scan(&tempf)

		var tempc float32
		tempc = conversao(tempf)

		fmt.Println(tempf, "FAHRENHEIT EQUIVALE A", tempc, "CELSIUS")
	}

}
