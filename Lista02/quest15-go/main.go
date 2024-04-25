package main

import (
	"fmt"
)

//JOSÉ

func main() {

	fmt.Println("DIGITE A QUANTIDADE E OS PARES DE NÚMEROS A SEREM COMPARADOS")

	var (
		fila1, fila2 []int
	)
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {

		var (
			num1, num2 int
		)

		fmt.Scan(&num1, &num2)

		fila1 = append(fila1, num1)

		fila2 = append(fila2, num2)

	}
	for j := 0; j < N; j++ {

		centena1 := fila1[j] / 100
		dezena1 := (fila1[j] / 10) % 10
		unidade1 := (fila1[j] % 10)
		centena2 := fila2[j] / 100
		dezena2 := (fila2[j] / 10) % 10
		unidade2 := (fila2[j] % 10)

		algarismo1 := unidade1*100 + dezena1*10 + centena1*1
		algarismo2 := unidade2*100 + dezena2*10 + centena2*1

		if algarismo1 > algarismo2 {

			fmt.Printf("%d\n", algarismo1)

		} else if algarismo1 < algarismo2 {

			fmt.Printf("%d\n", algarismo2)

		}

	}
}
