package main

import "fmt"

const nMAX = 100

//7 19 3 2 78 3 1 -3 18 19 0

func main() {
	var suara [nMAX]int
	var masuk, sah int
	var calon, jumlah int

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
		if jumlah > 0 { //Saat suara calon ke-i = 0, tidak perlu di print
			fmt.Println(calon, ":", jumlah)
			jumlah = 0
			calon++
		} else { //Saat suara calon ke-i > 0 perlu di print
			jumlah = 0
			calon++
		}
	}

}
