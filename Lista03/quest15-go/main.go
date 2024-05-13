package main

//MENOR DISTÂNCIA

import "fmt"

func menor_distancia(x []int) (int, int) {

	menordist := 2001

	var k int

	for i := 0; i < len(x); i++ {

		for j := i + 1; j < len(x); j++ {

			k = x[i] - x[j]

			if k < 0 {

				k = k * -1
			}
			if k < menordist {

				menordist = k
			}
		}
	}
	fator := len(x)

	combinacao := (fator * (fator - 1)) / 2

	return menordist, combinacao
}
func main() {

	fmt.Println("DIGITE")

	var testes int

	var seq2 []int

	fmt.Scan(&testes)

	for i := 0; i < testes; i++ {

		var num int

		fmt.Scan(&num)

		seq1 := make([]int, num)

		for i := range seq1 {

			fmt.Scan(&seq1[i])
		}
		a, b := menor_distancia(seq1)

		seq2 = append(seq2, a, b)
	}
	for i := 0; i < len(seq2); i += 2 {

		fmt.Printf("\n%v %v", seq2[i], seq2[i+1])
	}
}

//A resposta para a indagação da questão consiste no recurso, advindo da análise
// combinatória, "combinação". No exercício, é necessária a combinação de todos os
// números da sequência dada para realizar o cálculo da distância e encontrar a
// menor delas. Neste sentido, o primeiro número será combinado com todos os outros
// e, em seguida, o próximo número será combinado com todos os outros exceto o primeiro
// e assim sucessivamente. A quantidade de combinações cria, desta maneira, uma
// progressão aritimética que apresenta sempre um número maior que o outro e a soma
// destes valores representa a quantidade total de combinações.;)
