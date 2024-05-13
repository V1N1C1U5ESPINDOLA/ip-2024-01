package main

//ENTRADA INFINITA

import "fmt"

func main() {
	var nums []int

	fmt.Println("DIGITE OS NÚMEROS ( AO TERMINAR, DIGITE 0):")

	for {
		var num int

		fmt.Scan(&num)

		if num == 0 {
			break
		}

		nums = append(nums, num)
	}

	fmt.Println("NÚMEROS DIGITADOS:", nums)
}
