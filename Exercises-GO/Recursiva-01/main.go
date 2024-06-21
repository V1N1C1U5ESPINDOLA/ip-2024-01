package main

import "fmt"

func potencia(x, y int)(int){
	if y == 1{
		 return x
	}
	return x*potencia(x, y-1)
}

func main(){
	fmt.Println("DIGITE OS NÚMEROS")
	var num1,num2 int
	fmt.Scan(&num1,&num2)
	fmt.Printf("%v",potencia(num1,num2))
}