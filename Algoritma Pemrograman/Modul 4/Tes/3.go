package main

import "fmt"
import "math/rand"


func main() {
	var seed int64
	var hasil, tebakuser int

	fmt.Print("Angka Rahasia : ")
	fmt.Scanln(&seed)

	fmt.Print("Anda : ")
	fmt.Scanln(&tebakuser)

	rand.Seed(seed)
	hasildadang := rand.Intn(6)+1
	fmt.Println("Dadang :", hasildadang)
	
	hasil = rand.Intn(6)+1

	if 	tebakuser == hasildadang && tebakuser == hasil {
		fmt.Println("Nilai dadu", hasil, ", Seri")
	} else if hasildadang == hasil {
		fmt.Println("Nilai dadu", hasil, ", Pemenang adalah Dadang")
	} else if tebakuser == hasil {
		fmt.Println("Nilai dadu", hasil, ", Pemenang adalah Anda")
	} else {
		fmt.Println("Nilai dadu", hasil, ", Tidak ada pemenang")
	}
}