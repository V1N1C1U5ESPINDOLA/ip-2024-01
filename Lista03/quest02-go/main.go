package main

//CONTAGEM

import "fmt"

func main() {
	var (
		num, temp, K int
	)
	var (
		seq1, seq2 []int
	)
	fmt.Println("DIGITE OS NÚMEROS")

	for {
		fmt.Scan(&num)
		if num >= 1 && num <= 1000 {
			break
		}
	}

	for i := 0; i < num; i++ {

		fmt.Scan(&temp)
		seq1 = append(seq1, temp)
	}
	fmt.Scan(&K)

	for _, num2 := range seq1 {

		if num2 >= K {
			seq2 = append(seq2, num2)
		}
	}

	fmt.Printf("%v", len(seq2))

}
