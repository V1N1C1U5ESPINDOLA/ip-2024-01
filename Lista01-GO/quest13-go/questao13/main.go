package main

//CONVERSÃO DE NOTA EM CONCEITO

import "fmt"

func main() {

	fmt.Println("DIGITE SUA NOTA DE 0 À 10")
	var nota float32
	fmt.Scan(&nota)
	var conceito string

	if nota >= 9 {
		conceito = "A"
	} else if nota >= 7.5 && nota < 9.0 {
		conceito = "B"
	} else if nota >= 6.0 && nota < 7.5 {
		conceito = "C"
	} else if nota >= 0.0 && nota < 6.0 {
		conceito = "D"
	}
	fmt.Printf("NOTA %f\nCONCEITO %s", nota, conceito)
}
