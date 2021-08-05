package main

import "fmt"

const NMAX = 100

func main() {
	var n, x int
	var tab [NMAX]int
	var min1, min2 int
	var max1, max2 int

	fmt.Scanln(&n, &x) //Input n data dan x angka yang dicari

	fillTab(&tab, n) //Isi array

	//Cek x yang dicari ganjil/genap
	if x%2 == 0 { //Jika x genap
		findMin(tab, n, &min1, &min2)
		fmt.Println(min2, min1)
	} else { //Jika x ganjil
		findMax(tab, n, &max1, &max2)
		fmt.Println(max1, max2)
	}

}

func fillTab(tab *[NMAX]int, n int) { //Isi array
	i := 0
	for i < n-1 {
		fmt.Scan(&tab[i])
		i++
	}
}

func findMax(tab [NMAX]int, n int, max1, max2 *int) {
	*max1 = tab[0] //Index pertama dianggap max1 yang merupakan nilai maksimum sebenarnya
	for i := 0; i < n-1; i++ {
		if tab[i] > *max1 { //Jika tab[i] > max1, artinya tab[i] lebih besar dari max sebenarnya
			*max2 = *max1  //max1 yang lebih kecil dimasukkan nilainya ke variabel max2 yang merupakan nilai maksimum kedua
			*max1 = tab[i] //tab[i] yang nilainya lebih besar dari max sebenarnya, dimasukkan nilanya ke variable max1 menjadi max sebenarnya
		}
	}
}

func findMin(tab [NMAX]int, n int, min1, min2 *int) {
	*min1 = tab[0] //Index pertama dianggap min1 yang merupakan nilai minimum sebenarnya
	for i := 0; i < n-1; i++ {
		if tab[i] < *min1 { //Jika tab[i] < max1, artinya tab[i] lebih kecil dari min sebenarnya
			*min2 = *min1  //min1 yang lebih kecil dimasukkan nilainya ke variabel min2 yang merupakan nilai minimum kedua
			*min1 = tab[i] //tab[i] yang nilainya lebih kecil dari min sebenarnya, dimasukkan nilanya ke variable min1 menjadi min sebenarnya
		}
	}

}

/*
18 3
12 3 45 67 8 90 1 23 42 56 7 89 30 126 14 5 679 90

18 100
12 3 45 67 8 90 679 23 42 56 7 89 30 126 14 5 2 90

*/
