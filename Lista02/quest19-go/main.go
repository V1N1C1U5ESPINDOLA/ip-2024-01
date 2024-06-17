package main

// N AO CUBO
import (
	"fmt"
)

func main() {
	var m, p int
	fmt.Scanln(&m)

	for n := 1; n <= m; n++ {
		fmt.Printf("%d*%d*%d = ", n, n, n)
		p = (n * n) - (n - 1)
		fmt.Printf("%v", p)
		for i := 1; i < n; i++ {
			fmt.Printf("+%v", p+i*2)
		}
		fmt.Printf("\n")
	}

}
