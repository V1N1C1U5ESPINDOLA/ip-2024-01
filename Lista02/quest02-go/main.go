package main

//ULTRAPASSAGEM POPULACIONAL

import "fmt"

func main() {
	var (
		numpa, numpb, ano int
	)
	fmt.Println("DIGITE O NUMERO DE HABITANTES DO PAIS A")
	fmt.Scan(&numpa)
	fmt.Println("DIGITE O NUMERO DE HABITANTES DO PAIS B")
	fmt.Scan(&numpb)

	ano = 0
	for numpa <= numpb {
		numpa = int(float64(numpa) * 1.030)
		numpb = int(float64(numpb) * 1.015)
		ano++
	}
	fmt.Println(ano)
}
