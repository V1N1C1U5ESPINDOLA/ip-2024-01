package main

//CUSTO DA LATA DE CERVEJA

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("DIGITE O RAIO")
	var raio float32
	fmt.Scan(&raio)

	fmt.Println("DIGITE A ALTURA")
	var altura float32
	fmt.Scan(&altura)

	var custo float32

	custo = ((2 * (math.Pi * raio * raio)) + (2 * math.Pi * raio * altura)) * 100

	fmt.Println("O VALOR DO CUSTO É =", " ", custo)
}
