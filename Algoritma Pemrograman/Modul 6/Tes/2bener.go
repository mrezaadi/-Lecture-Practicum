package main

import "fmt"

func main() {
	var n int
	var m1, m2, m3 string

	total := 0 //varible total merupakan varible temporary dan akan terus berganti setiap kali fungsi dijalankan

	fmt.Scanln(&n, &m1, &m2, &m3) // meminta nilai dari variable m1, m2, m3

	if m1 == "f" {
		fungsif(n, &total)
	} else if m1 == "g" {
		fungsig(n, &total)
	} else if m1 == "h" {
		fungsih(n, &total)
	}

	//pengkondisian diatas bertujuan untuk mengeksekusi fungsi sesuai inputan dari varible m1,
	//dengan memasukkan nilai dari varible n dan akan mengubah varible total sebagai outputnya

	if m2 == "f" {
		fungsif(total, &total)
	} else if m2 == "g" {
		fungsig(total, &total)
	} else if m2 == "h" {
		fungsih(total, &total)
	}
	// pengkondisian diatas bertujuan untuk mengeksekusi fungsi yang selanjutnya akan dipakai
	// fungsi yang dipakai ditentukan dari nilai m2
	// dengan memasukkan keluaran dari fungsi pertama (var total)
	// dan akan mengubah nilai dari varible total tersebut dengan nilai hasil dri fungsi yg dijalankan

	if m3 == "f" {
		fungsif(total, &total)
	} else if m3 == "g" {
		fungsig(total, &total)
	} else if m3 == "h" {
		fungsih(total, &total)
	}

	// pengkondisian diatas bertujuan untuk mengeksekusi fungsi yang selanjutnya akan dipakai
	// fungsi yang dipakai ditentukan dari nilai m3
	// dengan memasukkan keluaran dari fungsi kedua (var total)
	// dan akan mengubah nilai dari varible total tersebut dengan nilai hasil dri fungsi yg dijalankan

	fmt.Println(total) //mengularkan hasil dari dijalankannya kombinasi ketiga fungsi diata
}

func fungsif(n int, total *int) {
	*total = 2*n + 5
	//menggunakan pointer untuk mengubah nilai total
}
func fungsig(n int, total *int) {
	*total = (n * n) + (2 * n)
	//menggunakan pointer untuk mengubah nilai total
}
func fungsih(n int, total *int) {
	*total = n - 3
	//menggunakan pointer untuk mengubah nilai total
}