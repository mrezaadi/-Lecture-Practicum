package main

import "fmt"



func main(){

	var total, usaha, nDoa int
	var doaOrtu bool
	var nilai string

	BacaData(usaha, nDoa, doaOrtu, nilai)

	TabungUsahaDoa(usaha, nDoa, &total)

	total = TabungDoaOrtu(doaOrtu, total)

	HasilNilaiAlpro(nilai, &total)

	fmt.Println(HasilAkhir(total))
}

func BacaData(usaha,jmlDoa int, doaOrtu bool, nilai string) {
    fmt.Print("Banyaknya usaha? ") ; fmt.Scanln(&usaha)
    fmt.Print("Banyaknya doa? ") ; fmt.Scanln(&jmlDoa)
    fmt.Print("Doa orang tua? ") ; fmt.Scanln(&doaOrtu)
    fmt.Print("Nilai Algoritma? ") ; fmt.Scanln(&nilai)
}

func TabungUsahaDoa(usaha, doa int, total *int){
	*total = usaha + doa
}

func TabungDoaOrtu(doa bool, total int) int{
	if doa {
		return total * 2
	} else {
		return total
	}
}

func HasilNilaiAlpro(nilai string, total *int){
	if nilai == "A"{
		*total = *total - 150
	} else if nilai == "B"{
		*total = *total - 130
	} else if nilai == "C"{
		*total = *total - 100
	}
}

func HasilAkhir(poin int) string{
	if poin >= 130{
		return "Lulus langsung dapat kerja gaji 2 digit"
	} else if poin >= 50 {
		return "Lulus langsung dapat kerja"
	} else {
		return "Jangan lelah berdoa dan berusaha, tidak ada yang sia-sia dari usaha dan doa"
	}
}
