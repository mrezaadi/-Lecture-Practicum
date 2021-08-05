package main

import "fmt"

func main(){
	var budget, jmlhwst, jmlhhr int
	var bT,bP,bK int
	var totalB, totalK int

	fmt.Print("Masukkan budget dalam dolar: ")
	fmt.Scanln(&budget) //Minta input budget

	fmt.Print("Masukkan jumlah tempat wisata dikunjungi: ")
	fmt.Scanln(&jmlhwst) //Minta input jumlah tempat wisata

	fmt.Print("Lama menginap: ")
	fmt.Scanln(&jmlhhr) //Minta input lama menginap

	hitungBiayaSatuan(jmlhhr, jmlhwst, &bT, &bP, &bK) //Menghitung biaya tiket masuk wisata, penginapan, dan konsumsi
	hitungTotalBiaya(bT, bP, bK, &totalB) //Menghitung total biaya tiket masuk wisata, penginapan, konsumsi, dan biaya perjalanan pulang pergi
	totalK = budget - totalB //Total kekurang atau kelebihan adalah budget - total biaya

	fmt.Println("Biaya tiket masuk: $ ", bT)
	fmt.Println("Biaya penginapan : $ ", bP)
	fmt.Println("Biaya konsumis: $ ", bK)

	fmt.Println("Total biaya perjalanan: $ ", totalB, "(Rp. ", UStoIDR(totalB),")")
	fmt.Println("Total kekurangan/kelebihan: $ ", totalK, "(Rp. ", UStoIDR(totalK),")")


}

func UStoIDR(dolar int)int{
	return dolar*13000 //1 dolar = Rp. 13.000
}

func hitungBiayaSatuan(jmlhhr, jmlhwst int, tiket, penginapan, konsumsi *int){
	*tiket = jmlhwst*13 //Biaya 1 tempat wisata $13
	*penginapan = jmlhhr*514 //Biaya 1 hari menginap $514 
	*konsumsi = jmlhhr*(2*20) //Biaya 1 kali makan $20, dan 1 hari makan 2 kali 
}

func hitungTotalBiaya(tiket, penginapan, konsumsi int, total *int){
	*total = tiket + penginapan + konsumsi + 450 //Total biaya tiket masuk wisata, penginapan, konsumsi, dan biaya perjalanan pulang pergi 
}