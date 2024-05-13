package main

//LER E IMPRIMIR

import "fmt"

func main() {
	var num, temp int

	fmt.Println("DIGITE OS NÚMEROS")
	fmt.Scan(&num)
	var seq1 []int

	for i := 0; i < num; i++ {

		fmt.Scan(&temp)

		seq1 = append(seq1, temp)

	}
	fmt.Print(seq1)
}
