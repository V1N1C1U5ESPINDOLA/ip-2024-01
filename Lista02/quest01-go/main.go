package main

//TURMA DE INTRODUÇÃO À PROGRAMAÇÃO

import "fmt"

func main() {
	var (
		np1, np2, np3, np4, np5, np6, np7, np8, nl1, nl2, nl3, nl4, nl5, nt, presenca float64
	)
	var matricula int
	fmt.Println("DIGITE SUA MATRICULA, SUAS NOTAS E FREQUENCIA")
	fmt.Scan(&matricula, &np1, &np2, &np3, &np4, &np5, &np6, &np7, &np8, &nl1, &nl2, &nl3, &nl4, &nl5, &nt, presenca)

	var (
		mediaprova, medialista, notafinal float64
	)

	mediaprova = (np1 + np2 + np3 + np4 + np5 + np6 + np7 + np8) / 8
	medialista = (nl1 + nl2 + nl3 + nl4 + nl5) / 5
	notafinal = (0.7*mediaprova + 0.15*medialista + 0.15*nt)

	var situacao string

	if presenca < 96.0 && notafinal < 6.0 {

		situacao = "REPROVADO POR NOTA E POR FREQUENCIA"

	} else if notafinal >= 6.0 && presenca < 96.0 {

		situacao = "REPROVADO POR FREQUENCIA"

	} else if notafinal < 6.0 && presenca >= 96.0 {

		situacao = "REPROVADO POR NOTA"

	} else {

		situacao = "APROVADO"
	}
	fmt.Println("MATRICULA:",matricula,"NOTA FINAL",notafinal,"SITUACAO FINAL",situacao)
}
