package main

import "fmt"

func main(){
	var winner, player, temp string
	var i, ronde int
	var nilai, answer int

	winner = "A"
	player = "B"
	ronde = 1

	fmt.Println("Ronde", ronde, ": ")
	fmt.Print(winner, " - masukkan angka rahasia: ")
	fmt.Scanln(&nilai)

	for nilai != -101{
		i = 1
		fmt.Print(player, " -masukkan angka tebakan ke- ", i,": ")
		fmt.Scanln(&answer)
		for i < 3 && answer != nilai {
			i = i + 1
			fmt.Print(player,"- masukkan angka tebakan ke-", i,": ")
			fmt.Scanln(&answer)
		}
		if nilai == answer {
			temp = winner
			winner = player
			player = temp
		}
		fmt.Println(winner,"adalah pemenangnya")
		ronde = ronde + 1
		fmt.Println("Ronde", ronde, ": ")
		fmt.Print(winner,"- masukkan angka rahasia: ")
		fmt.Scanln(&nilai)
	}
	fmt.Println("Permainan selesai")
}