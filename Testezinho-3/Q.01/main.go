package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func roda(x string) string {
	size := len(x) - 1
	seq2 := []rune(x)
	var seq1 []rune
	for i := 0; i <= size; i++ {
		seq1 = append(seq1, seq2[size-i])
	}
	return string(seq1)
}

func main() {
	fmt.Println("DIGITE AS FRASES INVERTIDAS")
	var seq []string
	for {
		var texto string
		ler := bufio.NewReader(os.Stdin)
		texto, _ = ler.ReadString('\n')
		texto = strings.TrimSpace(texto)
		if texto == "#" {
			fmt.Println("FIM DE ENTRADA")
			break
		}
		seq = append(seq, roda(texto))
	}
	for i := range seq {
		fmt.Printf("%v\n", seq[i])
	}
}
