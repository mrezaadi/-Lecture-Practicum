package main

import "fmt"

func main(){
	var n int
	var i int
	var a, b int
	var kunci int

	i = 0
	fmt.Scanln(&n)
	for i < n {
		fmt.Scanln(&a, &b)
		if ((a+b) % 2) == 0 {
			kunci = a
			fmt.Println(kunci,"\n")
		} else if ((a+b) % 2) == 1 {
			kunci = b
			fmt.Println(kunci,"\n")
		}
		i++
	} 
}