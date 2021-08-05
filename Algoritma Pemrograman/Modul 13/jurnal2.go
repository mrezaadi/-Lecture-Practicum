package main

import "fmt"

func main() {
	var tab [100]int
	var n, d, u int

	fmt.Scanln(&n, &u, &d) //Input banyak data , index start u, index stop d
	isiArray(&tab, n)      //Procedure isi array
	sorting(&tab, u, d, n) //Urutkan data secara descending, yang dimulai dari index ke u sampai index ke d
	tampilArray(tab, n)    //Ouptukan array yang sudah diurutkan
}

func isiArray(tab *[100]int, n int) {
	for i := 0; i < n; i++ { //Isi array dari index ke-0  sampai index ke-n
		fmt.Scan(&tab[i]) //Masukkan data angka
	}

}

func tampilArray(tab [100]int, n int) {
	for i := 0; i < n; i++ { //Menampilkan array dari index ke-0 sampai index ke-n
		fmt.Print(tab[i], " ") //Output setiap anggota array dipisahkan spasi
	}
}

func sorting(tab *[100]int, u, d, n int) { //Pengurutan dengan algoritma insertion sort secara descending dari index ke-u sampai index-d
	var pass, i, temp int

	pass = u + 1   //Dimulai dari index ke-u
	for pass < d { //Sampai index ke-d
		i = pass
		temp = tab[pass]
		for i > u && temp > tab[i-1] { //Cari posisi
			tab[i] = tab[i-1]
			i--
		}
		tab[i] = temp //Penyisipan
		pass++
	}
}

// 18 5 14
// 26 36 69 84 65 53 84 78 25 5 16 37 87 38 10 25 73 41
