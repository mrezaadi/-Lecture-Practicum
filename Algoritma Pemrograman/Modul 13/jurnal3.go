package main

import "fmt"

const NMAX = 100

type belanja [NMAX]int

func main() {
	var t belanja
	var n, ha int
	var hp float64

	entri(&t, &n)    //Masukkan data
	ha = total(t, n) //Menghitung total belanjaan awal

	if ha > 100000 { //Cek kondisi apakah total belanja lebih dari 100000
		urut(&t, n)         //Urutkan total belanja dari total belanja suatu barang yang paling kecil
		promo(t, n, &hp)    //Hitung harga setelah diberikan promo
		fmt.Println(ha, hp) //Outpukan harga setelah promo

	} else {
		fmt.Println(ha, ha) //Jika kurang dari 100000 maka harga tetap sama
	}

}

func entri(t *belanja, n *int) {
	var harga, kuantitas int
	var total int

	for i := 0; i < NMAX; i++ {
		fmt.Scanln(&harga, &kuantitas)    //Inputkan harga barang dan kuantitas
		if harga != 0 && kuantitas != 0 { //Cek apakah input trigger berhenti atau bukan
			total = harga * kuantitas //Harga yang dimasukkan ke dalam array adalah harga yang sudah dikalikan kuantitasnya
			t[i] = total              //Harga yang sudah dikalikan kuantitasnya dimasukkan ke dalam array
			*n++                      //Banyak jenis barang yang diinputkaan
		} else { //Jika inputan merupakan trigger maka perulangan berhenti
			break
		}
	}
}

func urut(t *belanja, n int) { //Pengurutan dengan algoritma insertion sort secara ascending
	var pass, i, temp int

	pass = 1
	for pass < n {
		i = pass
		temp = t[pass]
		for i > 0 && temp < t[i-1] { //Cari posisi
			t[i] = t[i-1]
			i--
		}
		t[i] = temp //Penyisipan
		pass++
	}

}

func total(t belanja, n int) int {
	var total int

	for i := 0; i < n; i++ {
		total = total + t[i] //Menghitung total belanjaan dengan cara menjumlah setiap elemen array dari index ke-0 sampai ke-n
	}
	return total
}

func promo(t belanja, n int, hp *float64) {
	var promo int            //Untuk menyimpan total belanjaan setelah dikenakan promo
	for i := 0; i < n; i++ { //Menghitung promo yang diberikan dari index k-0 sampai ke-n
		promo = promo + ((t[i]) - ((t[i] * (i + 1)) / 100)) //Promo yang diberikan dimulai dari 1% untuk barang termurah pertama, sedangkan barang termurah ke-n akan mendapat diskon sebesar n%
	}
	*hp = float64(promo) //Total belanjaan yang ada dalam variable promo, dimasukkan ke dalam variabel hp
}

// 16800 2
// 12100 3
// 9100 1
// 18300 2
// 13300 1
// 11300 3
// 0 0
