package main

import "fmt"

func main(){
	var n1, n2 int //Input angka
	var besar, kecil int //Perbandingan terbesar dan terkecil
	var kondisi bool //Kondisi pertama

	kondisi = true
	

	fmt.Scan(&n1, &n2)
	if n2 <= 0{ //Saat kondisi angka kedua tidak memenuhi
		besar = n1
		kecil = n1
	} else{
		for n2 > 0{ //Saat kondisi angka kedua memenuhi, dan terus minta input sampai kondisi kedua tidak memenuhi
			if kondisi == true{ //Kondisi pertama
				besar = max(n1,n2)
				kecil = min(n1,n2)
				kondisi = false
				fmt.Scan(&n2)
			} else if n2 >= besar{ //Bandingan terbesar
				besar = n2
				fmt.Scan(&n2)
			} else if n2 <= kecil{//Bandingan terkecil
				kecil = n2
				fmt.Scan(&n2)
			} else { //Saat tidak > terbsar atau < dari terkecil
				fmt.Scan(&n2)
			}
		}
	}
	fmt.Println(besar, kecil)
}

func max(f1,f2 int) int{
	var terbesar int

	if f1 > f2{ //Cek terbesar dari 2 angka
		terbesar = f1
	} else {
		terbesar = f2
	}
	return terbesar
}

func min(f1,f2 int) int{
	var terkecil int
	if f1 < f2 { //Cek terkecil dari 2 angka
		terkecil = f1
	} else {
		terkecil = f2
	}
	return terkecil
}