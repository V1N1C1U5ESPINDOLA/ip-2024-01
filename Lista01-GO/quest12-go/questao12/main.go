package main

//LOCADORA DE CHARRETES

import "fmt"

func main() {

	fmt.Println("DIGITE O TEMPO DE LOCAÇÃO")
	var tempo int
	fmt.Scan(&tempo)

	var (
		IN, RES int
	)
	IN = tempo / 3
	RES = tempo % 3

	var valor int
	valor = (IN * 10) + (RES * 5)

	fmt.Println("O VALOR A PAGAR É", " ", valor)

}
