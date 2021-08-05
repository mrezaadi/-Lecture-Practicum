package main

import "fmt"

func main() {
	var n, m, hasil, i, j int

	fmt.Scanln(&n)
	i = 0
	hasil = 0

	for i < n {
		fmt.Scanln(&m)
		if m > -1 && m < 10 {
			j = 0
			for j < i {
				m = m * 10
				j++
			}
			hasil = hasil + m
			i++
		}
	}
	fmt.Println(hasil)
}