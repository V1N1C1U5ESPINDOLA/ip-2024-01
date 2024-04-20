package main

//FATORIAL

import "fmt"

func main() {

	fmt.Println("DIGITE UM NÚMERO INTEIRO")
	var (
		num, i, f int
	)
	fmt.Scan(&num)
	f = 1
	if num == 0 {
		fmt.Println(f)
	} else {

		for i = 1; i <= num; i++ {

			f = f * i
		}

		fmt.Println(num, "! =", f)

	}

}
