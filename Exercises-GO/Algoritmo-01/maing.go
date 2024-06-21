package main

import "fmt"

//FUNÇÃO ORDENADORA
func BubbleSort(lista []int) []int {
	size := len(lista) - 1
	for i := 0; i < size; i++ {
		for j := size; j > i; j-- {
			if lista[j-1] > lista[j] {
				lista[j-1], lista[j] = lista[j], lista[j-1]
			}
		}
	}
	return lista
}

//FUNÇÃO BUSCADORA
func BinarySearch(lista []int, procurado int) int {
	maior := len(lista) - 1
	menor := 0
	for menor <= maior {
		meio := (maior + menor) / 2
		if lista[meio] == procurado {
			return meio
		}
		if lista[meio] > procurado {
			maior = meio - 1
		}
		if lista[meio] < procurado {
			menor = meio + 1
		}
	}
	return -1
}

func main() {
	var num int
	var seq []int
	var seek int
	fmt.Println("DIGITE O TAMANHO DA LISTA")
	fmt.Scan(&num)
	fmt.Println("DIGITE OS NÚMEROS DA LISTA")
	for i := 0; i < num; i++ {
		var temp int
		fmt.Scan(&temp)
		seq = append(seq, temp)
	}
	fmt.Println("DIGITE O NÚMERO A SER PROCURADO")
	fmt.Scan(&seek)
	fmt.Println("ORIGINAL :")
	for i := 0; i < len(seq); i++ {
		fmt.Printf("%v ", seq[i])
	}
	fmt.Println("\nORDENADA :")
	seq = BubbleSort(seq)
	for j := 0; j < len(seq); j++ {
		fmt.Printf("%v ", seq[j])
	}
	Bin := BinarySearch(seq, seek)
	fmt.Printf("\nPROCURADO : \nIndice=%v\n%v° Posição", Bin, Bin+1)
}
