package main

import "fmt"

func main(){
	var suara int 
	var s1,s2,s3 int

	//Nilai default suara, suara calon 1, suara calon 2, dan suara calon 3
	suara  = 0
	s1 = 0
	s2 = 0
	s3 = 0

	for suara != -1{ //Saat input tidak -1(mark), maka looping dilakukan
		fmt.Scan(&suara) //Minta input suara
		if suara == 1 || suara == 2 || suara == 3{ //Suara yang valid hanya 1, 2, dan  3
			hitungSuara(suara, &s1, &s2, &s3) //Panggil prosedur hitungSuara
		} 
	}

	fmt.Println(s1,s2,s3)
	
}

func hitungSuara(suara int, calon1, calon2, calon3 *int){
	if suara == 1{ //Saat input 1, suara akan ditambahkan ke total suara calon 1
		*calon1++
	} else if suara == 2{ //Saat input 2, suara akan ditambahkan ke total suara calon 3
		*calon2++
	} else if suara == 3{ //Saat input 3, suara akan ditambahkan ke total suara calon 3
		*calon3++
	}
}