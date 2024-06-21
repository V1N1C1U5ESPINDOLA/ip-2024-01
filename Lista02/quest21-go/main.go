package main

import "fmt"

func desordenadamente(x []float64) string {
	sus := true
	for i := 1; i < len(x); i++ {
		if x[i-1] > x[i] {
			sus = false
			break
		}
	}
	if sus == true {
		return "ORDENADA"
	}
	return "DESORDENADA"
}

func main() {
	var seq2 []string
	fmt.Println("DIGITE OS NÚMEROS")

	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		seq1 := make([]float64, num)
		for i := range seq1 {
			fmt.Scan(&seq1[i])
		}
		fmt.Printf("%v\n", seq1)
		result := desordenadamente(seq1)
		seq2 = append(seq2, result)
	}
	for i := range seq2 {
		fmt.Printf("%s\n", seq2[i])
	}
}
