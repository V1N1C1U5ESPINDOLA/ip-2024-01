package main

//LOTERIA

import "fmt"

func main() {

	var SENA, QUINA, QUADRA, contador int

	valor := map[int]int{

		6: SENA,
		5: QUINA,
		4: QUADRA,
	}
	premio := map[int]string{

		6: "sena",
		5: "quina",
		4: "quadra",
	}
	fmt.Println("DIGITE OS NÚMEROS")

	lucky := make([]int, 6)

	for i := range lucky {

		fmt.Scan(&lucky[i])
	}
	var num int

	fmt.Scan(&num)

	for i := 0; i < num; i++ {

		contador = 0

		slice := make([]int, 6)

		for j := range slice {

			fmt.Scan(&slice[j])
		}
		for k := 0; k < 6; k++ {

			for l := 0; l < 6; l++ {

				if slice[l] == lucky[k] {

					contador++

					break
				}
			}
		}
		switch contador {

		case 6:
			valor[6]++
		case 5:
			valor[5]++
		case 4:
			valor[4]++
		}
	}
	fmt.Printf("\n%v %v %v", valor[6], valor[5], valor[4])

	for i := 6; i > 3; i-- {

		if valor[i] == 0 {

			fmt.Printf("\nNenhum ganhador da %v", premio[i])

		} else {

			fmt.Printf("\n%v ganhadores da %v", valor[i], premio[i])
		}
	}
}
