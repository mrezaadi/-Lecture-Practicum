package main

import "fmt"

func main(){
	var n int //Bulan
	var data float64 //Jumlah nasabah
	var data1, data2 float64 //Perbadingan nilai untuk cek penambahan
	var jumlah_nasabah float64 //Total jumlah nasabah
	var selisih float64 //Selisih saat terjadi penambahan
	var penambahan float64 //Total selish penambahan
	var jumlah_penambahan float64 //Total terjadi penambahan
	var rataan_nasabah float64 //Rata-rata nasabah perbulan
	var rataan_penambahan float64 //Rata-rata terjadi penambahan


	n = 1
	jumlah_nasabah = 0
	data1 = 0

	for n <= 12 { //Perulangan 12 bulan
		fmt.Print("Bulan ", n,": ")
		fmt.Scanln(&data)
		if data < 0 { //Saat kondisi tidak memenuhi
		} else if data >= 0{ //Saat kondisi memenuhi
			jumlah_nasabah = jumlah_nasabah + data
			if data1 == 0{ //Untuk data pertama
				data1 = data
				n++
			} else if data1 != 0{ //Untuk data selanjutnya
				data2 = data
				if data2 > data1{ //Untuk terjadi penambahan
					selisih = data2-data1
					penambahan = penambahan + selisih
					jumlah_penambahan = jumlah_penambahan + 1
					data2, data1 = data1, data2 //Swap variabel
					n++
				} else if data2 < data1{ //Tidak terjadi penambahan
					data2, data1 = data1, data2  //Swap variabel
					n++
				}
			}

		}
	}
	rataan_nasabah = jumlah_nasabah/12.0
	rataan_penambahan = penambahan/jumlah_penambahan
	fmt.Println("Rataan nasabah :", rataan_nasabah)
	fmt.Println("Rataan penambahan :", rataan_penambahan)
}