package main

import "fmt"

type saco struct {
	prof string

	aula string

	carga int

	greve bool
}

func main() {
	var seq1 []saco

	fmt.Println("DIGITE, RESPECTIVAMENTE O NOME A MATERIA E A CARGA HORÁRIA\nDIGITE INF PARA FINALIZAR")

	for {
		var nome, materia string
		fmt.Scan(&nome)

		if nome == "INF" {
			break
		}

		fmt.Scan(&materia)

		var hora int
		fmt.Scan(&hora)

		var status bool
		fmt.Scan(&status)

		temp := saco{

			prof: nome,

			aula: materia,

			carga: hora,

			greve: status,
		}
		seq1 = append(seq1, temp)

	}

	for _, y := range seq1 {

		var s string

		if y.greve == true {

			s = "Sim"

		} else {

			s = "Não"

		}

		fmt.Printf("\nProfessor(a): %s\nDiciplina: %s\nCarga horária: %v\nAderiu à greve?: %s\n", y.prof, y.aula, y.carga, s)
	}

}
