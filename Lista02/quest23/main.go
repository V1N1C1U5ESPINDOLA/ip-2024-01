package main

//PROCURA POR UM NÚMERO AMIGO

import "fmt"

func somatorio(x int) int {
	soma := 0
	for i := 1; i <= int(x/2); i++ {
		if x%i == 0 {
			soma += i
		}
	}
	return soma
}

func main() {
	fmt.Println("DIGITE A QUANTIDADE DE NÚMEROS AMIGOS")
	var num int
	fmt.Scan(&num)
	if num > 9 {
		fmt.Println("NÚMERO GRANDE DEMAIS\nquer travar o pc (°~°) ?")
	} else {
		cont := 0
		for j := 1; cont < num; j++ {
			k := somatorio(j)
			if k > j && somatorio(k) == j {
				fmt.Printf("%v,%v\n", j, k)
				cont++
			}
		}
	}
}
