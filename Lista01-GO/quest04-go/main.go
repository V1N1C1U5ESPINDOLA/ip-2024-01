package main

//CONSUMO DE ENERGIA

import "fmt"

func main() {
	fmt.Println("DIGITE O SALÁRIO MÍNIMO ATUAL")

	var salario float32
	fmt.Scanln(&salario)

	fmt.Println("DIGITE A QUANTIDADE DE kW GASTA PELA SUA RESIDÊNCIA")

	var gasto float32
	fmt.Scanln(&gasto)

	var precokw float32
	precokw = (7 * salario) / 1000

	var precoconsumo float32
	precoconsumo = (gasto * precokw)

	var precodesconto float32
	precodesconto = (precoconsumo - (precoconsumo / 10))

	fmt.Println("Custo por kW:", " ", precokw)

	fmt.Println("Custo do consumo:", " ", precoconsumo)

	fmt.Println("Custo com desconto", " ", precodesconto)

}
