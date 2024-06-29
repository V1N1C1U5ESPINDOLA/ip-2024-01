package main

import "fmt"

func trocaOpostosSeMenor(x []int, n int) {
	for i := 0; i < int(n/2); i++ {
		oposto := n - 1 - i
		if x[i] < x[oposto] {
			troca(x, i, oposto)
		}
	}
}

func troca(seq3 []int, i, j int) {
	seq3[i], seq3[j] = seq3[j], seq3[i]
}

func main() {
	fmt.Println("DIGITE OS CASOS DE TESTE, O TAMANHO E OS NÚMEROS DA LISTA")
	var casos int
	fmt.Scan(&casos)
	for i := 0; i < casos; i++ {
		var size int
		fmt.Scan(&size)
		seq4 := make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scan(&seq4[j])
		}
		trocaOpostosSeMenor(seq4, size)
		for _, N := range seq4 {

			fmt.Printf("%v ",N)
		}
		fmt.Printf("\n")
	}
}
