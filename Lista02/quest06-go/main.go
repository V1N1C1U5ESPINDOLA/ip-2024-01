package main

//MAIOR SEGMENTO IGUAL DE UMA SEQUÊNCIA

import "fmt"

func main() {

	fmt.Println("DIGITE A QUANTIDADE DE NÚMEROS DA SEQUÊNCIA, EM SEGUIDA DIGITE OS COMPONENTES DESTA.")
	var num int
	fmt.Scan(&num)
	sequencia := make([]int, num)

	for i := 0; i < num; i++ {

		fmt.Scan(&sequencia[i])

	}
	var (
		sequenciaigual, sequenciamaioratual int
	)

	for j := 1; j < num; j++ {

		if sequencia[j] == sequencia[j-1] {
			sequenciaigual++

		} else {
			if sequenciaigual > sequenciamaioratual {

				sequenciamaioratual = sequenciaigual
			}

			sequenciaigual = 0
		}

	}
	if sequenciaigual > sequenciamaioratual {

		sequenciamaioratual = sequenciaigual
	}

	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos", sequenciamaioratual+1)
}
