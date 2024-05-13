package main

//ACHEI

import "fmt"

func seek(x []int, y int) bool {

	for _, z := range x {

		if z == y {
			return true
		}
	}
	return false

}

func main() {

	var (
		N, M, J, P int
	)
	var (
		V, S []int
	)

	fmt.Println("DIGITE OS NÚMEROS")

	fmt.Scan(&N)

	for i := 0; i < N; i++ {

		fmt.Scan(&P)
		V = append(V, P)

	}

	fmt.Scan(&M)

	for i := 0; i < M; i++ {

		fmt.Scan(&J)

		if seek(V, J) == true {

			S = append(S, 1)
		} else {
			S = append(S, 0)
		}

	}
	for i := 0; i < len(S); i++ {
		switch S[i] {
		case 1:
			fmt.Println("ACHEI")
		case 0:
			fmt.Println("NÃO ACHEI")
		}
	}
}
