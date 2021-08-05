package main

import "fmt"

var a,b,c,d int

func main() {
	fmt.Scanln(&a, &b, &c, &d)

	if d > a && d > b && d > c{
		fmt.Println("Ada rekor baru")
	} else {
		fmt.Println("Tidak ada rekor baru")
	}
}