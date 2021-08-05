package main

import "fmt"

const nMAX = 1000000

var data [nMAX]int

func main() {
	var n, k int
	var index int
	/*
		12 534
		1 3 8 16 32 123 323 323 534 543 823 999

		12 535
		1 3 8 16 32 123 323 323 534 543 823 999
	*/

	fmt.Scanln(&n, &k)   //Input n dan k
	isiArray(n)          //Isi array data
	index = posisi(n, k) //Cari posisi k di dalam array data

	if index == -1 {
		fmt.Println("TIDAK ADA") //Kondisi saat k tidak ditemukan
	} else {
		fmt.Println(index) //Kondisi saat k ditemukan
	}

}

func isiArray(n int) {
	i := 0
	for i < n {
		fmt.Scan(&data[i]) //Input array data
		i++
	}
}

func posisi(n, k int) int {
	var index int

	i := 0
	for i < n {
		if k == data[i] { //Jika ditemukan maka index == i
			index = i
			break
		} else { //Jika tidak ditemukan maka index == -1
			index = -1
			i++
		}
	}
	return index
}
