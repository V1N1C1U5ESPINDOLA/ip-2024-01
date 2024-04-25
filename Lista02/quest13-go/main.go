package main

//GRÃOS DE MILHO NO TABULEIRO DE XADREZ

import "fmt"

func main() {

	fmt.Println("DIGITE UM NÚMERO INTEIRO")
	var (
		num, multi int
	)
	fmt.Scan(&num)

	multi = ((num * 96) - num)

	fmt.Printf("%d\n", multi)

}
