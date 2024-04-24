package main

import "fmt"

func main() {

	fmt.Printf("DIGITE O VALOR DO INGRESSO\nDIGITE O VALOR INICIAL\nDIGITE O VALOR FINAL\n")
	var (
		valoringresso, valorinicial, valorfinal, difmenor, difmaior float64
	)

	fmt.Scan(&valoringresso, &valorinicial, &valorfinal)

	var quantidade []float64

	for i := valorinicial; i <= valorfinal; i++ {

		quantidade = append(quantidade, i)
	}
	for j := 0; j < len(quantidade); j++ {

		if quantidade[j] > valoringresso {

			difmaior = (quantidade[j] - valoringresso)
			ingressos := (120 - (difmaior/0.50)*30)
			lucro := (ingressos * quantidade[j]) - 200 - (0.05 * ingressos)
			fmt.Printf("V:%.2f,N:%.2f,L:%.2f\n", quantidade[j], ingressos, lucro)

		} else if quantidade[j] < valoringresso {

			difmenor = (valoringresso - quantidade[j])
			ingressos := (120 + (difmenor/0.50)*25)
			lucro := (ingressos * quantidade[j]) - 200 - (0.05 * ingressos)
			fmt.Printf("V:%.2f,N:%.2f,L:%.2f\n", quantidade[j], ingressos, lucro)

		}

	}
}
