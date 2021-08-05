package main

import "fmt"

func main() {
	var x int
	fmt.Scanln(&x)

	if x%3 == 0{
		fmt.Println("Fizz")
	} 
	if x%5 == 0{
		fmt.Println("Bazz")
	}
}