package main

//ARRECADAÇÃO DE JOGOS

import (
	"fmt"
)

func main() {
	var numerodecasos int32

	fmt.Println("DIGITE O NÚMERO DE CASOS")
	fmt.Scanln(&numerodecasos)
	var i int32
	for i = 1; i <= numerodecasos; i++ {

		var Totalpublico float32
		var popular float32
		var geral float32
		var arquibancada float32
		var cadeiras float32
		fmt.Println("DIGITE O TOTAL DO PUBLICO")
		fmt.Scanln(&Totalpublico)
		fmt.Println("DIGITE O PUBLICO POPULAR")
		fmt.Scanln(&popular)
		fmt.Println("DIGITE O PUBLICO GERAL")
		fmt.Scanln(&geral)
		fmt.Println("DIGITE O PUBLICO ARQUIBANCADA")
		fmt.Scanln(&arquibancada)
		fmt.Println("DIGITE O PUBLICO CADEIRAS")
		fmt.Scanln(&cadeiras)

		var Npopular float32
		var Ngeral float32
		var Narquibancada float32
		var Ncadeiras float32
		var RENDATOTAL float32

		Npopular = (Totalpublico / 100) * popular
		Ngeral = (Totalpublico / 100) * geral
		Narquibancada = (Totalpublico / 100) * arquibancada
		Ncadeiras = (Totalpublico / 100) * cadeiras

		RENDATOTAL = Npopular*1 + Ngeral*5 + Narquibancada*10 + Ncadeiras*20

		fmt.Println("A RENDA TOTAL DO JOGO É \n", RENDATOTAL)

	}
}
