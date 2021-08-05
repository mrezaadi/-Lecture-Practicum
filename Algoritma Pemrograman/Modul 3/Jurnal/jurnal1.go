package main

import "fmt"

func main(){

	var a, b, c, d,e int
	var f, g, h  rune

	fmt.Scanln(&a, &b, &c, &d, &e)
	fmt.Scanf("%c%c%c", &f, &g, &h)

	fmt.Printf("%c%c%c%c%c\n", a, b-32,  c-32, d-32, e-32)
	fmt.Printf("%c%c%c", f+1, g+1, h+1)

}