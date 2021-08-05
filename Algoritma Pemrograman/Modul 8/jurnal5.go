package main

import "fmt"

func main(){
	var p1,p2 string
	var soal1, skor1, soal2, skor2 int
	
	//Nilai default jumlah soal dan skor dari peserta1 dan peserta2
	soal1 = 0
	soal2 = 0
	skor1 = 0
	skor2 = 0


	i := 1

	for i <= 2{ //Perulangan untuk minta input sebanyak dua kali
		if i == 1 {
			fmt.Scan(&p1) //Input nama peserta1
			hitungSkor(&soal1,&skor1) //Panggil prosedur hitung skor
			i++ //Iterasi
		} else{
			fmt.Scan(&p2) //Input nama peserta2
			hitungSkor(&soal2, &skor2) //Panggil prosedur hitung skor
			fmt.Println("Selesai")
			i++ //Iterasi
		}
	}

	if soal1 == soal2{ //Saat jumlah soal yang dikerjakan sama, masuk ke kondisi berikutnya
		if skor1<skor2{ //Jika lama waktu pengerjaan(skor) peserta1 < lama waktu pengerjaan(skor) peserta 2. Maka peserta 1 pemenangnya
			fmt.Println(p1,soal1,skor1)
		} else{ //Jika lama waktu pengerjaan(skor) peserta1 > lama waktu pengerjaan(skor) peserta 2. Maka peserta 2 pemenangnya
			fmt.Println(p2,soal2,skor2)
		}
	} else if soal1>soal2{ //Jika jumlah soal yang dikerjakan peserta1 > peserta2. Maka peserta1 pemenangnya
		fmt.Println(p1, soal1, skor1)
	} else{ //Jika jumlah soal yang dikerjakan peserta1 < peserta2. Maka peserta2 pemenangnya
		fmt.Println(p2,soal2,skor2)
	}

}

func hitungSkor(soal, skor *int){
	var waktu int //Lama peserta mengerjakan soal, yang dijadikan skor dalam penilaian
	j := 1 
	for j <= 8{ //Perulangan sebanyak 8 kali
		fmt.Scan(&waktu) //Input waktu mengerjakan 1 soal (skor)
		if waktu >= 301{ //Saat soal tidak dikerjakan, waktu otomatis 301, jumlah soal yang diekerjakan peserta tidak dihitung, lama waktu tidak dihitung sebagai skor
			j++ //Iterasi
		} else{ //Saat kondisi selain kondisi diatas
			j++ //Iterasi
			*soal++ //Jumlah soal yang dikerjakan bertambah 1
			*skor = *skor + waktu //Skor bertambah seusai dengan waktu yang diinputkan
		}
	}
}