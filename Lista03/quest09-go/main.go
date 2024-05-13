package main

//DISTÂNCIA ENTRE PONTOS 3D

import (
	"fmt"
	"math"
)

func distance(x []float64) []float64 {

	var dist float64

	var seq2 []float64

	for i := 0; i < (len(x) - 3); i += 3 {

		dist = (math.Pow((x[i] - x[i+3]), 2)) + (math.Pow((x[i+1] - x[i+4]), 2)) + (math.Pow((x[i+2] - x[i+5]), 2))

		seq2 = append(seq2, math.Sqrt(dist))

	}
	return seq2
}
func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var (
		num, temp float64
	)

	var (
		seq1, seq3 []float64
	)
	fmt.Scan(&num)

	if num < 2 {

		fmt.Println("NÚMERO INVÁLIDO")

	}
	for j := 0.0; j < num; j++ {

		for i := 0.0; i < 3; i++ {

			fmt.Scan(&temp)

			seq1 = append(seq1, temp)
		}
	}
	seq3 = distance(seq1)

	for i := 0; i < len(seq3); i++ {
		fmt.Printf("\n%.3v", seq3[i])
	}
}
