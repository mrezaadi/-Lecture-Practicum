package main

import "fmt"

func main() {
	var text string
	fmt.Scanln(&text)

	for text != "selesai"{
		fmt.Println(text)
		fmt.Scanln(&text)
	}
}