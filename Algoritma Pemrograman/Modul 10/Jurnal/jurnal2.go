package main

import "fmt"

const MAX = 107

/*
9 130440 130939 130210 130879 130793 130943 130531 130823 130879
6 130939 130531 130440 130141 130943 130879

1 130461
1 130652

9 13062 130593 130542 13027 13015 130172 130475 130996 13091
9 13062 130593 130542 13027 13015 130172 130475 130996 13091
*/

func main() {
	var n, m int
	var mahasiswa [MAX]string
	var wisuda [MAX]string

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mahasiswa[i]) //Scan array mahasiswa
	}

	fmt.Scan(&m)
	for j := 0; j < m; j++ {
		fmt.Scan(&wisuda[j]) //Scan array wisuda
	}

	updateKelulusan(&mahasiswa, wisuda, &n, m) //Updat kelulusan
	fmt.Print(n, " ")                          //Print sisa jumlah mahasiswa
	for k := 0; k < n; k++ {
		fmt.Print(mahasiswa[k], " ") //Print NIM mahasiswa yang tersisa
	}

}

func posisi(tab [MAX]string, n int, x string) int {
	var index int

	index = -1

	for i := 0; i < n; i++ { //Cari apakah x ada di dalam array tab
		if x == tab[i] {
			index = i
			break //Saat ditemukan retrun indexnya
		}
	}
	return index //Saat tidak ditemukan index akan -1
}

func hapus(tab *[MAX]string, n *int, x string) {
	var index int
	index = posisi(*tab, *n, x)
	if index != -1 { //Saat index = -1, maka tidak terjadi penghapusan elemen
		for i := index; i < *n; i++ {
			tab[i] = tab[i+1] //Hapus index yang menunjukkan nilai x dengan cara menggeser seluruh elemen di depannya
		}
		*n = *n - 1 //Setiap terjadi penghapusan maka jumlah mahasiswa berkurang 1
	}

}

func updateKelulusan(mahasiswa *[MAX]string, wisuda [MAX]string, n *int, m int) {
	i := 0
	for i < m {
		hapus(mahasiswa, n, wisuda[i]) //Cek semua elemen pada aray wisuda apakah ada di array mahasiswa, jika ada maka elemen tersebut dihapus
		i++
	}
}
