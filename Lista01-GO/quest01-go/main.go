package main

//APROVADO OU REPROVADO

import (
	"fmt"
)

func main() {

	var nota01 float32
	var nota02 float32
	var nota03 float32
	var soma float32
	var media float32

	fmt.Println("DIGITE SUAS NOTAS ESCOLARES")
	fmt.Scanln(&nota01, &nota02, &nota03)
	soma = nota01 + nota02 + nota03
	media = soma / 3

	fmt.Println(media)
	var mensagem string

	if media < 6 {
		mensagem = "REPROVADO"
	} else {
		mensagem = "APROVADO"
	}
	fmt.Print("Média Escolar:", " ", media, "\n")
	fmt.Print("Futuro:", " ", mensagem, "\n")
}
