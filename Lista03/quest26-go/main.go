package main

import "fmt"

func fake_nanismo(x []int) []int {

	x = BubbleSort(x)

	var seq3 []int
	var soma int

	for _, val := range x {

		soma += val
	}
	resto := soma - 100

	for i := 0; i < 8; i++ {

		for j := i + 1; j < 9; j++ {

			if x[i]+x[j] == resto {

				for k := 0; k < 9; k++ {

					if k != i && k != j {

						seq3 = append(seq3, x[k])
					}
				}
				return seq3
			}
		}
	}
	return seq3
}

func BubbleSort(l []int) []int {

	n := len(l)

	for i := 0; i < n-1; i++ {

		for j := n - 1; j > i; j-- {

			if l[j-1] > l[j] {

				l[j-1], l[j] = l[j], l[j-1]
			}
		}
	}
	return l
}

func main() {

	var num int

	fmt.Scan(&num)

	var seq []int

	for i := 0; i < num*9; i++ {

		var val int

		fmt.Scan(&val)

		seq = append(seq, val)
	}
	for j := 0; j < num; j++ {

		var seq1 []int

		for i := 0; i < 9; i++ {

			seq1 = append(seq1, seq[j*9+i])
		}
		seq2 := fake_nanismo(seq1)

		for _, k := range seq2 {

			fmt.Println(k)
		}
	}
}
