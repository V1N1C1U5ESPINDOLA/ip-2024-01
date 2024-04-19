package main

//TEMPO EM SEGUNDOS

import "fmt"

func main() {

	fmt.Println("DIGITE AS HORAS, MINUTOS E SEGUNDOS")
	var (
		H, M, S int
	)
	fmt.Scan(&H, &M, &S)
	TS := H*60*60 + M*60 + S

	fmt.Println("O TEMPO EM SEGUNDOS É", TS)
}
