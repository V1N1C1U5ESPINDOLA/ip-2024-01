package main

import "fmt"

func inverte(x []int, y int) []int {

	if y == 0 {
		x = x[len(x)/2:]
		return x
	} else {
		x = append(x[:], x[y-1:y]...)
		return inverte(x, y-1)
	}
}

func main() {
	fmt.Print("DIGITE A QUANTIDADE E OS NÚMEROS DA SEQUÊNCIA\n")
	var num int
	fmt.Scan(&num)
	seq := make([]int, num)
	for i := 0; i < num; i++ {
		var temp int
		fmt.Scan(&temp)
		seq[i] = temp
	}
	fmt.Printf("%v", inverte(seq, num))
	fmt.Printf("\n%v",seq)
}
