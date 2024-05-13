package main

//CONTAGEM DE ELEMENTOS ÚNICOS

import "fmt"

func elemental(x []int) int {

	var unico int

	freq := make(map[int]int)

	for _, i := range x {

		freq[i]++
	}
	for _, j := range freq {

		if j == 1 {

			unico++
		}
	}
	return unico
}
func main() {

	fmt.Println("DIGITE OS NÚMEROS")

	var num, temp int

	fmt.Scan(&num)

	var seq1 []int

	for i := 0; i < num; i++ {

		fmt.Scan(&temp)

		seq1 = append(seq1, temp)
	}
	fmt.Printf("%v", elemental(seq1))
}
