package main

//QUADRADO DE PARES

import "fmt"

func main() {

	fmt.Println("DIGITE UM VALOR INTEIRO ENTRE 5 E 2000")
	var num int
	fmt.Scan(&num)
	var i int
	for i = 2; i <= num; i += 2 {
		fmt.Println(i, "^ 2 =", i*i)
	}

}
