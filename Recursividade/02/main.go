package main

import "fmt"

func soma(x []float64) []float64 {
	size := len(x) - 1
	if size == 0 {
		return x
	} else {
		x[0] = x[0] + x[size]
		x = append(x[:size], x[size+1:]...)
	}
	return soma(x)
}

func main() {
	fmt.Println("DIGITE O TAMANHO E A OS NÚMEROS DA SEQUÊNCIA")
	var seq []float64
	var num int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		var temp float64
		fmt.Scan(&temp)
		seq = append(seq, temp)
	}
	fmt.Printf("%v", soma(seq))
}
