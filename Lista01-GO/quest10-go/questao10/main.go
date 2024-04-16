package main

//CALCULO DO DETERMINANTE DE UMA MATRIZ QUADRADA DE DUAS DIMENSÕES

import "fmt"

func main() {

	fmt.Println("DIGITE OS VALORES DOS QUATRO ELEMENTOS DA MATRIZ")
	var (
		a, b, c, d, determinante float32
	)
	fmt.Scan(&a, &b, &c, &d)

	determinante = a*d - b*c

	fmt.Println("O VALOR DA DETERMINANTE É =", " ", determinante)

}
