package main

import "fmt"

const NMAX int = 1000000

type tabDate [NMAX]string

var t tabDate
var n int

func main() {
	addData(&t, &n)
	urutData(&t, n)
	printData(t, n)
}

func addData(list *tabDate, n *int) {
	var tanggal string

	for i := 0; i < NMAX; i++ {
		fmt.Scanln(&tanggal)
		if tanggal != "####.##.##" {
			list[i] = tanggal
			*n++
		} else {
			break
		}
	}
}

func urutData(list *tabDate, n int) {
	var idx int

	for i := 0; i < (n - 1); i++ {
		idx = maxPos(*list, i, n)
		swap(&list[i], &list[idx])
	}
}

func maxPos(list tabDate, start, n int) int {
	var max string
	var hasil int

	max = list[start]
	hasil = start
	start++
	for start < n {
		if list[start] > max {
			max = list[start]
			hasil = start
		}
		start++
	}
	return hasil
}

func swap(x, y *string) {
	*x, *y = *y, *x
}

func printData(list tabDate, n int) {
	fmt.Println("")
	for i := 0; i < n; i++ {
		fmt.Println(list[i])
	}
}
