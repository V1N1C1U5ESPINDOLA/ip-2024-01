package main

import "fmt"

type fundador struct {
	x, y, z float64
}

func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var num int

	fmt.Scan(&num)

	var seq1 []fundador

	var seq3 []float64

	for i := 0; i < num; i++ {

		var X, Y, Z float64

		fmt.Scan(&X, &Y, &Z)

		temp := fundador{

			x: X,
			y: Y,
			z: Z,
		}
		seq1 = append(seq1, temp)
	}
	for j := 0; j < (num - 1); j++ {

		seq2 := make([]float64, 3)

		seq2[0] = seq1[j+1].x - seq1[j].x
		seq2[1] = seq1[j+1].y - seq1[j].y
		seq2[2] = seq1[j+1].z - seq1[j].z

		maior := seq2[0]

		if maior < 0 {

			maior = -maior
		}
		for k := 1; k < 3; k++ {

			if seq2[k] < 0 {

				seq2[k] = -seq2[k]
			}
			if seq2[k] > maior {

				maior = seq2[k]
			}
		}
		seq3 = append(seq3, maior)
	}
	for _, valor := range seq3 {
		fmt.Printf("%.2f\n", valor)
	}
}
