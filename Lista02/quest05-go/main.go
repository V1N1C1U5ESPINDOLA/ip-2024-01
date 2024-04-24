package main

//MAIOR SEGMENTO CRESCENTE DE UMA SEQUÊNCIA

import (
	"fmt"
)

func main() {
	fmt.Println("DIGITE A QUANTIDADE DE NÚMEROS DA SEQUÊNCIA")
	var num int
	fmt.Scan(&num)

	sequencia := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Scan(&sequencia[i])
	}

	var maiorcomprimento int
	var comprimentoatual int
	comprimentoatual = 1

	for j := 1; j < num; j++ {
		if sequencia[j] > sequencia[j-1] {
			comprimentoatual++
		} else {
			if comprimentoatual > maiorcomprimento {
				maiorcomprimento = comprimentoatual
			}
			comprimentoatual = 1
		}
	}
	comprimentofinal := maiorcomprimento - 1
	fmt.Printf("O comprimento do segmento crescente máximo é: %d", comprimentofinal)
}
