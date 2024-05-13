package main

// MAIOR ELEMENTO

import "fmt"

func main() {

	fmt.Println("DIGITE OS NÚMEROS")
	var seqfinal []int

	for {

		var (
			num, temp, maioratual, maioridice int
		)
		var (
			seq1 []int
		)

		fmt.Scan(&num)

		if num == 0 {
			break
		}

		for i := 0; i < num; i++ {

			fmt.Scan(&temp)

			seq1 = append(seq1, temp)

		}
		for j := 0; j < num; j++ {

			if seq1[j] > maioratual {

				maioratual = seq1[j]
				maioridice = j

			}

		}

		seqfinal = append(seqfinal, maioridice, maioratual)

		for i := 0; i < len(seqfinal); i += 2 {
			fmt.Printf("%d %d\n", seqfinal[i], seqfinal[i+1])
		}
	}

}
