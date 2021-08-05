package main

import "fmt"

func main() {

	var r,t float64
	fmt.Scanln(&r, &t)
	fmt.Println(hitungVolume(r,t))
}

func hitungVolume(r float64, t float64) float64 {
	var volume float64
	volume = r * r * 3.14 * t
	return volume
}