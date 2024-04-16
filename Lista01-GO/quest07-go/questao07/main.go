package main

//CONVERSÃO PARA O SISTEMA MÉTRICO

import "fmt"

func main() {

	fmt.Println("DIGITE A TEMPERATURA EM FAHRENHEIT")
	var tempf float32
	fmt.Scan(&tempf)

	fmt.Println("DIGITE A QUANTIDADE DE CHUVA EM POLEGADA")
	var polegada float32
	fmt.Scan(&polegada)

	var tempc float32
	tempc = (5 * (tempf - 32)) / 9

	var chuva float32
	chuva = (polegada * 25.4)

	fmt.Println("O VALOR EM CELSIUS =", " ", tempc)
	fmt.Println("A QUANTIDADE DE CHUVA É =", " ", chuva)
}
