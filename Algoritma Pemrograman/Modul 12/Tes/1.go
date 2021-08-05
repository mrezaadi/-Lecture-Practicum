package main

import (
	"fmt"
	"math"
)

const NMAX = 100

type provinsi struct {
	nama     string
	populasi int
	tumbuh   float64
}

type data_provinsi struct {
	tabel [NMAX]provinsi
	n     int
}

func main() {
	var data data_provinsi
	var x1, x2 string
	var datax1 provinsi
	var prediksix2 int

	fmt.Println("Data Pertumbuhan Provinsi:")
	isiData(&data) //Panggil prosedur isiData untuk mengisi data

	fmt.Print("Nama provinsi? ")
	fmt.Scan(&x1)                                            //Input nama provinsi yang ingin ditampilkan datanya
	datax1 = cariProvinsi(data, x1)                          //Panggil fungsi cariProvinsi untuk medapatkan data provinsi yang dicari
	fmt.Println(datax1.nama, datax1.populasi, datax1.tumbuh) //Menampilakn data provinsi berupa nama, populasi, dat pertumbuhan penduduk

	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scan(&x2)                   //Input nama provinsi yang ingin dicari prediksi penduduknya untuk tahun depan
	prediksix2 = prediksi(data, x2) //Panggil fungsi prediksi untuk menghitung prediksi, dan mengembalikan hasil prediksinya
	fmt.Print(prediksix2)           //Menampilkan hasil prediksi

	fmt.Println("\nUrutan data pertumbuhan provinsi berdasarkan populasi terurut menurun:")
	urutData(&data)               //Panggil fungsi urutData untuk mengurutkan data secara terurut menurun atau descending
	for i := 0; i < data.n; i++ { //Menampilkan data yang sudah diurutkan secara descending
		fmt.Println(data.tabel[i].nama, data.tabel[i].populasi, data.tabel[i].tumbuh)
	}
}

func isiData(tab *data_provinsi) {
	var prov string
	var pop int
	var tumbuh float64

	for i := 0; i < NMAX; i++ {
		fmt.Scan(&prov, &pop, &tumbuh)
		if prov != "NONE" && pop != 0 && tumbuh != 0.0 { //Cek apakah input trigger berhenti atau bukan
			tab.tabel[i].nama = prov
			tab.tabel[i].populasi = pop
			tab.tabel[i].tumbuh = tumbuh
			tab.n++
		} else { //Jika input trigger berhenti, maka perulangan berhenti
			break
		}
	}
}

func cariProvinsi(tab data_provinsi, x string) provinsi {
	var found provinsi

	for i := 0; i < tab.n; i++ { //Cari stringx pada tab data_provinsi
		if tab.tabel[i].nama == x {
			found = tab.tabel[i]
			break
		}
	}
	return found
}

func prediksi(tab data_provinsi, x string) int {
	var prov provinsi
	var hasil_prediksi float64

	prov = cariProvinsi(tab, x) //Cari data provinsi yang mau dihitung prediksinya

	hasil_prediksi = float64(prov.populasi) + (float64(prov.populasi) * prov.tumbuh) //Menghitung prediksi

	return int(math.Ceil(hasil_prediksi)) //Sesuai contoh, hasil dari perhitungan prediksi akan dibulatkan ke atas atau ceiling
}

func urutData(tab *data_provinsi) {
	var pass, idx, i int
	var temp provinsi

	//Sorting dengan metode selection sort
	pass = 1
	for pass < tab.n {
		idx = pass - 1
		i = pass
		for i < tab.n {
			if tab.tabel[idx].populasi < tab.tabel[i].populasi {
				idx = i
			}
			i++
		}
		temp = tab.tabel[pass-1]
		tab.tabel[pass-1] = tab.tabel[idx]
		tab.tabel[idx] = temp
		pass++
	}
}

// Aceh 4906835 0.0201
// Bali 4104900 0.0133
// Banten 11704877 0.0131
// Bengkulu 1844800 0.0259
// Gorontalo 1115633 0.018
// Sumut 13766851 0.0261
// Yogyakarta 3553100 0.0195
// NONE 0 0.0
