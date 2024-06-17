package main

import (
	"fmt"
	"strconv"
)

func naobinarie(x int) string {
	if x == 0 {
		return "0"
	}
	if x == 1 {
		return "1"
	}
	return naobinarie(x/2) + strconv.Itoa(x%2)
}

func main() {
	fmt.Print("DIGITE O NÚMERO A SER CONVERTIDO:\n")
	var num int
	fmt.Scan(&num)
	fmt.Printf("%s\n", naobinarie(num))
}
