package main

import "fmt"

func main(){
	var day1, day2, year1, year2 int
	var day, month1, month2 string
	var m2 int

	fmt.Scanln(&day1, &month1, &year1, &day)

	Pengambilan(day1, Bulan2Angka(month1), year1, day, &day2, &m2, &year2)
	month2 = Angka2Bulan(m2)
	fmt.Println(day2, month2, year2)
}

func Kabisat(tahun int) bool {
	return tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
}

func Bulan2Angka(bulan string) int{
	var angka int
	if bulan == "Januari"{
		angka = 1
	} else if bulan == "Februari"{
		angka = 2
	} else if bulan == "Maret"{
		angka = 3
	} else if bulan == "April"{
		angka = 3
	} else if bulan == "Mei"{
		angka = 5
	} else if bulan == "Juni"{
		angka = 6
	} else if bulan == "Juli"{
		angka = 7
	} else if bulan == "Agustus"{
		angka = 8
	} else if bulan == "Spetember"{
		angka = 9
	} else if bulan == "Oktober"{
		angka = 10
	} else if bulan == "November"{
		angka = 11
	} else if bulan == "Desember"{
		angka = 12
	}
	return angka
}

func Angka2Bulan(angka int)string{
	var bulan string
	if angka == 1{
		bulan = "Januari"
	} else if angka == 2{
		bulan = "Februari"
	} else if angka == 3{
		bulan = "Maret"
	} else if angka == 4{
		bulan = "April"
	} else if angka == 5{
		bulan = "Mei"
	} else if angka == 6{
		bulan = "Juni"
	} else if angka == 7{
		bulan = "Juli"
	} else if angka == 8{
		bulan = "Agustus"
	} else if angka == 9{
		bulan = "Spetember"
	} else if angka == 10{
		bulan = "Oktober"
	} else if angka == 11{
		bulan = "November"
	} else if angka == 12{
		bulan = "Desember"
	}
	return bulan
}

func JumlahHari(bln, thn int) int{
	if bln == 1 || bln == 3 || bln == 5 || bln == 7 || bln == 8 || bln == 10 || bln == 12 {
		return 31
	} else if bln == 4 || bln == 6 || bln == 9 || bln == 11 {
		return 30
	} else {
		if Kabisat(thn) {
			return 29
		} else {
			return 28
		}
	}
}

func Pengambilan(tgl1,bln1,thn1 int, hari string, tgl2, bln2, thn2 *int){
	var waktu_proses int

	if hari == "Kamis" || hari == "Jumat"{
		waktu_proses = 4
	} else {
		waktu_proses = 2
	}
	*tgl2 = tgl1 + waktu_proses
	totalHari := JumlahHari(bln1, thn1)
	if *tgl2 <= totalHari{
		*bln2 = bln1
		*thn2 = thn1
	} else {
		*tgl2 = *tgl2 - totalHari
		if bln1== 12{
			*bln2 = 1
			*thn2 = thn1 + 1
		} else {
			*bln2 = bln1 + 1
			*thn2 = thn1
		}
	}
}