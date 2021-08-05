package main

import "fmt"

func main() {

	var nilai int
	var jumlah int
	var n int
	var rata2 float64

	fmt.Scanln(&nilai)

	for nilai != -1{
		n = n + 1
		jumlah = jumlah + nilai
		fmt.Scanln(&nilai)
	}

	if n == 0{
		rata2 = 0.0
	} else {
		rata2 = float64(jumlah)/float64(n)
	}

	fmt.Println(rata2)
}