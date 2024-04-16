package main

//DIVISÍVEL POR 3 E 5

import "fmt"

func main() {

	fmt.Println("DIGITE UM NÚMERO INTEIRO")
	var num int
	fmt.Scan(&num)
	if num%15 == 0 {

		fmt.Println("O NÚMERO É DIVISÍVEL")
	} else {

		fmt.Println("O NÚMERO NÃO É DIVISÍVEL")
	}

}
