package main

import "fmt"

func main() {
	var (
		numeros   int
		divPar    int
		divImpar  int
		somaPar   float64
		somaImpar float64
	)

	fmt.Println("DIGITE UMA SEQUÊNCIA DE NÚMEROS:")

	for {

		fmt.Scan(&numeros)

		if numeros == 0 {
			break
		}

		if numeros%2 == 0 {
			somaPar += float64(numeros)
			divPar++
		} else {
			somaImpar += float64(numeros)
			divImpar++
		}
	}

	mediaPar := somaPar / float64(divPar)
	mediaImpar := somaImpar / float64(divImpar)

	fmt.Printf("MÉDIA PAR = %.2f\n", mediaPar)
	fmt.Printf("MÉDIA ÍMPAR = %.2f\n", mediaImpar)
}
