package main

//MEDIANA

import "fmt"

func ordem_na_parada(x []float64) []float64 {
	var j []float64
	ponto := (len(x))

	for i := 0; i < (ponto - 1); i++ {

		for j := 0; j < ((ponto - 1) - i); j++ {

			if x[j] > x[j+1] {

				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	j = x

	return j
}
func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var num float64

	var mediana float64

	fmt.Scan(&num)

	seq1 := make([]float64, int(num))

	for i := range seq1 {

		fmt.Scan(&seq1[i])
	}
	seq1 = ordem_na_parada(seq1)

	if len(seq1)%2 == 0 {

		soma := (seq1[int(len(seq1)/2)-1]) + seq1[int(len(seq1)/2)]

		mediana = soma / 2

	} else {

		mediana = seq1[(len(seq1)/2)-1/2]
	}
	fmt.Printf("%v", mediana)
}
