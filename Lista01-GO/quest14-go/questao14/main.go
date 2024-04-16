package main

//VOLUME DA PIRÂMIDE DE BASE HEXAGONAL

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("DIGITE OS VALORES DA ALTURA E ARESTA")
	var (
		altura, aresta float64
	)
	fmt.Scan(&altura, &aresta)
	var (
		abase, volume float64
	)
	abase = (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2
	volume = (1.0 / 3.0) * abase * altura

	fmt.Println("O VOLUME DA PIRÂMIDE É =", volume, "METROS CÚBICOS")

}
