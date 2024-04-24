package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var matricula int
	var (
		horasTrabalhadas,
		valorHora float64
	)

	for {
		fmt.Scanf("%d %f %f\n", &matricula, &horasTrabalhadas, &valorHora)
		if matricula == 0 {
			break
		}
		salario := float64(horasTrabalhadas) * valorHora
		fmt.Printf("%d %.2f\n", matricula, salario)

		bufio.NewReader(os.Stdin).ReadByte()
	}
}
