package main

import "fmt"

var s string
var a,b int
var hasil int

func main() {
	fmt.Scanln(&s, &a, &b)

	hasil = a + b

	fmt.Println("Kata =", s)
	fmt.Println("Jumlah =", hasil)
}