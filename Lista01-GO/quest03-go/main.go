package main

//COMPOSIÇÃO INTEIRA
import "fmt"

func main() {
	fmt.Println("DIGITE TRÊS ALGARISMOS")
	var n1, n2, n3 string
	fmt.Scanln(&n1, &n2, &n3)
	fmt.Println(n1 + n2 + n3)
}
