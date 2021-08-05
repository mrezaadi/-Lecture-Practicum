package main

import "fmt"

type buku struct{ //Tipe data bentukan buku
	judul, penulis string
	tahun int
}

func main(){

	var kardus [5]buku //Kapasitas kardus 5 buku
	var atas int //Index buku paling atas
	var find string //Judul buku yang dicari
	
	atas = -1 //Kondisi saat kardus kosong

	i := 1
	for i <= 4{ //Sesuai soal, inputkan 4 buku
		tambahBuku(&kardus, &atas)
		i++
	}

	find = "C" //Sesuai soal, cari buku dengan judul "C", asumsikan ada buku denga judul "C" seperti contoh

	cariBuku(&kardus, &atas, find) //Cari buku dengan judul yang ada pada variabel find

	tambahBuku(&kardus, &atas) //Sesuai soal, tambahkan buku hingga kardus penuh

	find = "NULL" //Sesuai soal, cari buku dengan judul yang tidak ada pada kardus, asumsikan tidak ada buku dengan judul "NULL"
	cariBuku(&kardus, &atas, find)
}

func tambahBuku(kardus *[5]buku, atas *int){
	*atas = *atas + 1 //Index buku paling atas bertambah 1, karena buku bertambah 1
	fmt.Print("\nMasukkan judul : ")
	fmt.Scanln(&kardus[*atas].judul) //Judul buku
	fmt.Print("Masukkan penulis : ")
	fmt.Scanln(&kardus[*atas].penulis) //Penulis buku
	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scanln(&kardus[*atas].tahun) //Tahun terbit
}

func ambilBuku(kardus *[5]buku, atas *int, ambil *buku){
	*ambil = kardus[*atas] //Buku yang diambil adalah buku yang berada di index paling atas
	*atas = *atas - 1 //Index atas berkurang 1, karena buku berkurang 1
}

func cariBuku(kardus *[5]buku, atas *int, x string){
	var kondisi bool //Kondisi apakah buku yang di cari ada atau tidak
	var keluar int //Index yang menunjukkan buku yang dikeluarkan
	var ambil buku //Buku yang diambil

	keluar = *atas //Buku yang dikeluarkan adalah buku yang paling ats

	i := 0

	for i <= *atas{
		if x == kardus[keluar].judul{ //Jika buku pada posisi array kardus ke index keluar sama dengan buku yang dicari, maka buku ketemu sehingga kondisi true, dan perulangan berhenti
			kondisi = true
			break
		} else{ //Jika buku pada posisi array kardus ke index keluar tidak sama dengan buku yang dicari, maka buku tidak ketemu sehingga kondisi false
			kondisi = false
			ambilBuku(kardus, &keluar, &ambil) //Ambil buku pada array kardus index keluar
			fmt.Print(ambil.judul, " ") //Buku yang di keluarkan di ouputkan
			i++
		}
	}
	if kondisi == true{
		fmt.Println(" KETEMU") //Jika perulangan selesai dan kondisi true, artinya buku ketemu
	} else{
		fmt.Println(" TIDAK KETEMU") //Jika perulangan selesai dan kondisi false, artinya buku tidak ketemu
	}
}