package main

// CÁLCULO DO DELTA NA EQUAÇÃO DE BÁSKARA

import "fmt"

func main() {

	fmt.Println("DIGITE O VALOR DO COEFICIENTE A")
	var coa float32
	fmt.Scan(&coa)

	fmt.Println("DIGITE O VALOR DO COEFICIENTE B")
	var cob float32
	fmt.Scan(&cob)

	fmt.Println("DIGITE O VALOR DO COEFICIENTE C")
	var coc float32
	fmt.Scan(&coc)

	var delta float32
	delta = cob*cob + (-4 * coa * coc)

	fmt.Println("O VALOR DE DELTA É =", " ", delta)

}
