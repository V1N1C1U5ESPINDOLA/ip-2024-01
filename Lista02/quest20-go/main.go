package main

// NÚMERO PERFEITO

import "fmt"

func main() {

	fmt.Println("DIGITE O NÚMERO")

	var num, total int
	var seq1 []int
	var strings string

	fmt.Scan(&num)

	for i := 1; i < num; i++ {

		if num%i == 0 {

			seq1 = append(seq1, i)

			total += i
		}
	}
	if total == num {

		strings = "(NÚMERO É PERFEITO)"

	} else {

		strings = "(NÚMERO NÃO É PERFEITO)"
	}
	fmt.Printf("%v= 1 ", num)

	for j := 1; j < len(seq1); j++ {

		fmt.Printf(" + %v", seq1[j])
	}
	fmt.Printf(" = %v %s", total, strings)
}
