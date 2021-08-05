package main

import "fmt"

const NMAX = 127

func main(){
	var kata [NMAX]string //Array yang menerima input tiap huruf yang di input
	var n int //Banyak huruf yang di input

	fmt.Print("Teks : ")
	isiArray(&kata, &n) //Minta input huruf yang membentuk kata yang dipisahkan dengan spasi
	fmt.Println("Palindrom :", palindrom(&kata, &n)) //Cek nilai kebenaran apakah katda palindrom atau bukan
}

func isiArray(kalimat *[NMAX]string, n *int){
	var huruf string

	i := 0 
	for i <= len(kalimat){
		fmt.Scan(&huruf)
		if huruf == "."{ //Trigger ".", saat input titik perulangan berhenti
			break
		} else{
			kalimat[i] = huruf //Inputan kata di assign ke array kalimat pada index ke i
			*n = *n + 1 //Setiap memasukkan 1 huruf, jumlah huruf bertambah 1
			i++
		}
	}
}

func cetakArray(kalimat *[NMAX]string, n *int){
	i := 0
	for i < *n-1{
		fmt.Print(kalimat[i]) //Cetak array kalimat index ke i
		i++
	}
}

func balikanArray(kalimat *[NMAX]string, n *int){
	i := 0 //Menunjukaan index dari huruf pertama
	j := *n-1 //Menunjukan index dari huruf terakhir
	for i < j {
		kalimat[i], kalimat[j] = kalimat[j], kalimat[i] 
		//Membalikkan array dengan cara menukar posisi huruf pertama , dengan huruf terakhir, dst
		i++
		j--
	}
}

func palindrom(kalimat *[NMAX]string, n *int)bool{
	var newArr [NMAX]string //Untuk membandingkan array yang sudah dibalik dengan array yang belum dibalik

	newArr = *kalimat //Array baru disamakan dengan array lama 
	balikanArray(&newArr, n) //Array baru yang berisi array lama dibalik isinya

	return newArr == *kalimat //Nilai kebenaran apakah array baru yang sudah dibalik tetap sama dengan array lama
}