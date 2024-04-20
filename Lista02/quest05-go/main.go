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

	for i := 0; i <= num; i++ {
		fmt.Scan(&sequencia[i])
	}
}
