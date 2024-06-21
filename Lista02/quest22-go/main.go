package main

//TRANSFORMA DECIMAL EM FRAÇÃO

import (
    "fmt"
    "strconv"
)

func fracionando(x float64) string {
    var nume, deno string
    resto:=1.0
    for i := 1.0; resto > 0.0000000007; i++ {
        resto = (x * i - float64(int(x * i)))
        if resto < 0.0000000007 {
            nume = strconv.Itoa(int(x * i))
            deno = strconv.Itoa(int(i))
        }
    }
    fracao := nume + "/" + deno
    return fracao
}

func main() {
    fmt.Println("DIGITE UM NÚMERO REAL")
    var num float64
    fmt.Scan(&num)
    fmt.Printf("FRAÇÃO RESPECTIVA = %v\n", fracionando(num))
}
