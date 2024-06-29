package main

import "fmt"

func BOLHA(x[]int)[]int{
	size:= len(x)-1

	for l:=0;l<size;l++{
		for m:=size;m>l;m--{
			if x[m-1]>x[m]{
				x[m-1], x[m]= x[m], x[m-1]
			}
		}
	}
	return x
}

func main (){
	fmt.Println("DIGITE A QUANTIDADE E OS NÚMEROS DA LISTA")
	var num int
	fmt.Scan(&num)
	seq:=make([]int,num)
	for i:=range seq{
		fmt.Scan(&seq[i])
	}
	fmt.Printf("ORIGINAL= %v\n",seq)
	fmt.Printf("ORDENADA= %v",BOLHA(seq))
}