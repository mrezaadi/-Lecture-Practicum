package main

import "fmt"

func main(){
	var n int

	fmt.Scanln(&n)

	fmt.Println((n/1000)%10)
	fmt.Println((n/100)%10)
	fmt.Println((n/10)%10)
	fmt.Println(n%10)
	fmt.Println((n/100)%100)
	fmt.Println((n/10)%100)
	fmt.Println(n%100)
	fmt.Println(n/10)
	fmt.Println((n%1000))
	fmt.Println(n)
}