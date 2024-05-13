package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, IndiceMenor, d int
	resultados := []string{}
	menor := 9

	for {
		fmt.Scan(&d, &num)
		if d < 1 || num < d || num > 100000 {
			fmt.Println("Fora dos limites")
			return
		}

		if d == 0 {
			break
		}
		numString := strconv.Itoa(num)
		for i := 0; i < d; i++ {
			for i := range numString {
				x, _ := strconv.Atoi(string(numString[i]))
				if x < menor {
					menor, _ = strconv.Atoi(string(numString[i]))
					IndiceMenor = i
				}
			}
			numString = numString[:IndiceMenor] + numString[IndiceMenor+1:]
			menor = 9
		}
		resultados = append(resultados, numString)
	}
	for i := range resultados {
		fmt.Println(resultados[i])
	}
}
