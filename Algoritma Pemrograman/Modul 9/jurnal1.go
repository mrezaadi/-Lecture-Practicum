package main

import "fmt"

const NMAX = 100

func main(){

	var t1,t2,t3 [NMAX]int //Untuk menampung kemenangan tiap tahun dari 3 tim
	var n1,n2,n3 int //Untuk menghitung tahun tiap tim
	
	fmt.Print("Masukkan kemenangan tim 1 : ")
	inputData(&t1,&n1) //Input data tim 1
	fmt.Print("Masukkan kemenangan tim 2 : ")
	inputData(&t2,&n2) //Input data tim 2
	fmt.Print("Masukkan kemenangan tim 3 : ")
	inputData(&t3,&n3) //Input data tim 3

	fmt.Println("\nRataan tim 1 adalah :",rataan(t1,n1)) //Hitung rataan kemenangan tim 1
	fmt.Println("Rataan tim 2 adalah :",rataan(t2,n2)) //Hitung rataan kemenangan tim 2
	fmt.Println("Rataan tim 3 adalah :",rataan(t3,n3)) //Hitung rataan kemenangan tim 3

}

func inputData(team *[NMAX]int, n *int){
	var win int
	i := 0
	for i < len(team)-1{
		fmt.Scan(&win) //Kemenangan setiap tahun
		if win == -1{ //Trigger -1, saat -1 perulangan berhenti
			break
		} else {
			team[i] = win //Inputan win di assign ke aray kemenangan pada index ke i
			*n = *n + 1 //Setiap input kemenangan ke dalam array kemenangan, tahun bertambah 1
			i++
		}
	}
}

func rataan(team [NMAX]int, n int)float64{
	var mean float64
	var sum_win int

	i := 0

	for i < n{ //Menghitungan total kemenangan dalam suatu tim
		sum_win = sum_win + team[i]
		i++
	}
	mean = float64(sum_win)/float64(n) //Total kemenangan/total tahun
	return mean
}