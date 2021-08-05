package main

import "fmt"

func main(){
	var kapasitas int //Kapasitas keranjang
	var i int //Iterasi
	var barang string //Barang 
	var status string //y atau n
	var jmlh_n int //jumlah status n
	
	fmt.Scanln(&kapasitas)

	i = 1
	jmlh_n = 0

	for i <= kapasitas{ //Melakukan perulangan sampai kapasitas penuh
		fmt.Scanln(&barang, &status)
		if status == "y"{ //Saat barang penting kapasitas tambah 1
			i++
		} else if status == "n" { //Saat barang tidak penting jumlah tidak penting tambah 1 dan tidak tambah kapasitas
			jmlh_n++
		}
	}
	fmt.Println(jmlh_n)
}