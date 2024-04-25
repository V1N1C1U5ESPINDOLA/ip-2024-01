package main

//LUCRO DE MERCADORIAS

import (
	"fmt"
)

func calucro(c, v float64) float64 {
	var lucro float64
	lucro = ((v - c) / c * 100)

	return lucro

}

func main() {

	var (
		fila1, fila4 []int
	)
	var (
		fila2, fila3, lucromenor, lucroentre, lucromaior, sequencialu []float64
	)
	var (
		lucromelhor, vtcompras, vtvendas, vtlucro float64
	)
	var (
		indicecod1, indicecod2, maisvenda int
	)

	fmt.Println("DIGITE OS NÚMEROS, ESCREVA -1 PARA TERMINAR")
	for {

		var (
			codigo, numvend int
		)
		var (
			precocompra, precovenda float64
		)

		fmt.Scan(&codigo)

		if codigo == -1 {

			break
		}

		fmt.Scan(&precocompra)

		fmt.Scan(&precovenda)

		fmt.Scan(&numvend)

		fila1 = append(fila1, codigo)

		fila2 = append(fila2, precocompra)

		fila3 = append(fila3, precovenda)

		fila4 = append(fila4, numvend)

	}

	for i := 0; i < len(fila2); i++ {

		lucro := calucro(fila2[i], fila3[i])

		sequencialu = append(sequencialu, lucro)

		if lucro < 10 {

			lucromenor = append(lucromenor, lucro)

		} else if lucro <= 20 && lucro >= 10 {

			lucroentre = append(lucroentre, lucro)

		} else if lucro > 20 {

			lucromaior = append(lucromaior, lucro)

		}

	}

	for j := 0; j < len(sequencialu); j++ {

		if sequencialu[j] > lucromelhor {

			lucromelhor = sequencialu[j]

			indicecod1 = j

		}

	}

	for k := 0; k < len(fila4); k++ {

		vtcompras = vtcompras + (fila2[k] * float64(fila4[k]))

		vtvendas = vtvendas + (fila3[k] * float64(fila4[k]))

		if fila4[k] > maisvenda {

			maisvenda = fila4[k]

			indicecod2 = k

		}

	}

	fmt.Println("QUANTIDADE DE MERCADORIAS QUE GERARAM LUCRO MENOR QUE 10%:", len(lucromenor))

	fmt.Println("QUANTIDADE DE MERCADORIAS QUE GERARAM LUCRO MAIOR OU IGUAL A 10% & MENOR OU IGUAL A 20%:", len(lucroentre))

	fmt.Println("QUANTIDADE DE MERCADORIAS QUE GERARAM LUCRO MAIOR QUE 20%:", len(lucromaior))

	fmt.Println("CÓDIGO DA MERCADORIA QUE GEROU MAIOR LUCRO:", fila1[indicecod1])

	fmt.Println("CÓDIGO DA MERCADORIA MAIS VENDIDA:", fila1[indicecod2])

	vtlucro = calucro(vtcompras, vtvendas)

	fmt.Printf("VALOR TOTAL DE COMPRAS: %.2f, VALOR TOTAL DE VENDAS: %.2f E PERCENTUAL DE LUCRO TOTAL: %.2f", vtcompras, vtvendas, vtlucro)
}
