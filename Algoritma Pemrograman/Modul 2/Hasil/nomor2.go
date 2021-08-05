package main

import "fmt"

var r int
var luas float64

func main() {
	fmt.Scanln(&r)
	luas = 22.0/7.0 * float64(r)

	fmt.Println("Luas lingkaran dengan jari-jari =", r,"adalah", luas)
}