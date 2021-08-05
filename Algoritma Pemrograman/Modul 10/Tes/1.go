package main

import "fmt"

const nMAX = 100

func main() {
	var suara [nMAX]int
	var masuk, sah int
	var calon, jumlah int
	var ketua, wakil int
	var jumlah1, jumlah2 int
	var kondisi bool

	i := 0

	for i < len(suara) {
		fmt.Scan(&suara[i]) //Scan suara masuk
		if suara[i] == 0 {  //Trigger berhenti meminta input suara
			break
		} else if suara[i] >= 1 && suara[i] <= 20 { //Jika suara yang masuk sah, yaitu antara 1 dan 20 inklusif
			masuk++
			sah++
			i++
		} else { //Jika suara yang masuk tidak sah
			masuk++
			i++
		}
	}

	fmt.Println("Suara masuk :", masuk)
	fmt.Println("Suara sah :", sah)

	calon = 1

	for calon <= 20 { //Cek jumlah suara setiap calon dari 1-20
		j := 0
		for j < masuk { //Cari suara dan jumlah suara calon ke-i di dalam array suara
			if calon == suara[j] { //Saat calon ke-i ditemukan
				jumlah++
				j++
			} else { //Saat calon ke-i tidak ditemukan
				j++
			}
		}

		if kondisi == true { //Untuk cek apakah inputan pertama atau bukan
			ketua = calon    //Inputan pertama di asumsikan paling banyak
			jumlah1 = jumlah //Jumlah suara ketua berada pada variable jumlah1
			kondisi = false  //Kondisi false agar tidak masuk ke kondisi ini lagi
		} else {
			jumlah2 = jumlah        //Inputan setelah inputan pertama. Assign jumlah suara inputan ke variabel jumlah2, jumlah2 digunakan untuk menyimpan jumlah suara wakil
			if jumlah1 == jumlah2 { //Jika jumlah1 == jumlah2. Maka wakil adalah calon dengan urutan yang lebih besar
				wakil = calon
			} else if jumlah2 > jumlah1 { //Jika jumlah2 > wakil. Artinya inputan setelah inputan pertama lebih banyak suaranya dari pertama
				wakil = calon
				ketua, wakil = wakil, ketua        //Swap nilai pada variable ketua dan wakil
				jumlah1, jumlah = jumlah2, jumlah1 //Swap nilai pada jumlah1 dan jumlah2
			}
		}
		jumlah = 0 //Reset kembail jumlah
		calon++    //Cek calon selanjutnya
	}
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil Ketua :", wakil)
}
