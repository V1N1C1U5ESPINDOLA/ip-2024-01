package main

//CONTA DE AGUA

import "fmt"

func main() {
	fmt.Println("DIGITE O NÚMERO DA SUA CONTA DE ÁGUA")
	var numconta int32
	fmt.Scanln(&numconta)

	fmt.Println("DIGITE O CONSUMO DE ÁGUA")
	var consumoagua float32
	fmt.Scanln(&consumoagua)

	fmt.Println("DIGITE O TIPO DE CONSUMIDOR")
	var consumidor string
	fmt.Scanln(&consumidor)

	var contavalor float32

	if consumidor == "R" {
		contavalor = 5 + 0.05*consumoagua
	} else if consumidor == "C" {
		if consumoagua <= 80 {
			contavalor = 500
		} else if consumoagua > 80 {
			contavalor = 500 + 0.25*(consumoagua-80)
		}
	} else if consumidor == "I" {
		if consumoagua <= 100 {
			contavalor = 800
		} else if consumoagua > 100 {
			contavalor = 800 + 0.04*(consumoagua-100)
		}
	}

	fmt.Println("CONTA =", " ", numconta)
	fmt.Println("VALOR DA CONTA =", " ", contavalor)
}
