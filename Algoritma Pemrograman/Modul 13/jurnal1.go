package main

import "fmt"

const NMAX = 100

type mahasiswa struct { //Tipe data mahasiswa berisi inisial nama dan tinggi badan
	nama   string
	tinggi int
}

type data_mahsiswa [NMAX]mahasiswa //Array berisi data mahasiswa

func main() {
	var data data_mahsiswa
	var n int

	bacaData(&data, &n) //Procedure input n dan data ke dalam array
	urutData(&data, n)  //Procedure mengurutkan data dalam array secara ascending dengan algoritma selection sort
	fmt.Println("\n")
	cetakData(data, n) //Ouputkan data yang sudah terurut ascending

}

func bacaData(tab *data_mahsiswa, n *int) {
	fmt.Scanln(n)             //Input jumlah masukkan
	for i := 0; i < *n; i++ { //Isi array dari index ke-0  sampai index ke-n
		fmt.Scanln(&tab[i].nama, &tab[i].tinggi) //Masukkan data nama dan tinggi ke dalam array
	}
}

func cetakData(tab data_mahsiswa, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(tab[i].nama, tab[i].tinggi) //Cetak data dalam array dari index 0 sampai n
	}
}

func urutData(tab *data_mahsiswa, n int) { //Pengurutan dengan algoritma selection sort
	var temp mahasiswa
	var pass, idx, i int

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if tab[idx].tinggi > tab[i].tinggi { //Cari nilai paling besar
				idx = i
			}
			i++
		}
		temp = tab[pass-1] //Swap variable
		tab[pass-1] = tab[idx]
		tab[idx] = temp
		pass++
	}
}

// 10
// VR 169
// NP 157
// US 165
// QF 162
// NH 156
// BP 164
// TV 157
// RC 159
// MS 166
// TE 158
