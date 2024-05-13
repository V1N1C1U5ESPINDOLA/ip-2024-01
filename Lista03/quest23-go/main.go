package main

//COMPARAÇÃO DE TEXTOS

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func linguistica(x string) []int {

	apar := make([]int, 5)

	for _, j := range x {
		switch j {

		case 'a', 'â', 'ã', 'á', 'à':
			apar[0]++
		case 'e', 'ê', 'é':
			apar[1]++
		case 'i', 'í':
			apar[2]++
		case 'o', 'ô', 'õ', 'ó':
			apar[3]++
		case 'u', 'ú':
			apar[4]++
		}
	}
	return apar
}

func calculo1(x, y []int) float64 {

	var somatorio float64

	for i := 0; i < 5; i++ {

		somatorio += math.Pow(float64(x[i]-y[i]), 2.0)
	}
	return math.Sqrt(somatorio)
}

func main() {
	fmt.Println("DIGITE AS FRASES")

	var texto string

	ler := bufio.NewReader(os.Stdin)

	texto, _ = ler.ReadString('\n')

	texto = strings.TrimSpace(texto)

	div := strings.Split(strings.ToLower(texto), ";")

	if len(div) != 2 {

		fmt.Println("FORMATO INVÁLIDO !")

		return
	}
	aeiou1 := linguistica(div[0])
	aeiou2 := linguistica(div[1])

	for i := 0; i < 5; i++ {
		fmt.Printf("%v,", aeiou1[i])
	}
	fmt.Printf("\n")
	for j := 0; j < 5; j++ {
		fmt.Printf("%v,", aeiou2[j])
	}
	fmt.Printf("\n%.2f", calculo1(aeiou1, aeiou2))
}
