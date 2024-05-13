package main

//COMPARAÇÃO DE TEXTOS

import "fmt"

func megadavirada(x int, y string) []string {

	delet := len(y) - x

	var pins []byte
	var strings []string

	for i := 0; i < len(y); i++ {

		if len(pins) > 0 && pins[len(pins)-1] < y[i] && delet > 0 {

			pins = pins[:len(pins)-1]

			delet = delet - 1
		}
		pins = append(pins, y[i])
	}
	strings = append(strings, string(pins[:x]))

	return strings
}

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var seq1 []string
	var algarismos, mantidos int
	var numero string

	for {
		fmt.Scan(&algarismos, &mantidos)

		if algarismos == 0 && mantidos == 0 {
			break
		}
		fmt.Scan(&numero)

		seq1 = append(seq1, megadavirada(mantidos, numero)...)
	}
	for i := range seq1 {

		fmt.Printf("\n%v", seq1[i])
	}
}
