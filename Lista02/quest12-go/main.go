package main

import "fmt"

func main() {

	fmt.Printf("DIGITE O VALOR DO INGRESSO\nDIGITE O VALOR INICIAL\nDIGITE O VALOR FINAL\n")
	var (
		valoringresso, valorinicial, valorfinal, difmenor, difmaior float64
	)

	fmt.Scan(&valoringresso, &valorinicial, &valorfinal)

	var quantidade []float64

	var melhorvalor float64

	var guardindice int

	var (
		fila, fila2, fila3 []float64
	)

	for i := valorinicial; i <= valorfinal; i++ {

		quantidade = append(quantidade, i)
	}
	for j := 0; j < len(quantidade); j++ {

		if quantidade[j] > valoringresso {

			difmaior = (quantidade[j] - valoringresso)

			ingressos := (120 - (difmaior/0.50)*30)

			lucro := (ingressos * quantidade[j]) - 200 - (0.05 * ingressos)

			fmt.Printf("V:%.2f,N:%.2f,L:%.2f\n", quantidade[j], ingressos, lucro)

			fila = append(fila, lucro)

			fila2 = append(fila2, ingressos)

			fila3 = append(fila3, quantidade[j])

		} else if quantidade[j] < valoringresso {

			difmenor = (valoringresso - quantidade[j])

			ingressos := (120 + (difmenor/0.50)*25)

			lucro := (ingressos * quantidade[j]) - 200 - (0.05 * ingressos)

			fmt.Printf("V:%.2f,N:%.2f,L:%.2f\n", quantidade[j], ingressos, lucro)

			fila = append(fila, lucro)

			fila2 = append(fila2, ingressos)

			fila3 = append(fila3, quantidade[j])
		}

	}
	var k int

	for k = 0; k < len(fila); k++ {

		if fila[k] > melhorvalor {

			melhorvalor = fila[k]
			guardindice = k

		}

	}

	fmt.Printf("MELHOR VALOR FINAL:%.2f\nLUCRO:%.2f\nNÚMERO DE INGRESSOS:%.2f", fila3[guardindice], melhorvalor, fila2[guardindice])
}
