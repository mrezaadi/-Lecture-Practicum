package main

import "fmt"
import "math"

func main(){

	var cx,cy,x,y,r float64
	var kondisi bool

	fmt.Scanln(&cx,&cy,&r) //Pusat taman
	fmt.Scanln(&x, &y) //Pusat anda

	kondisi = posisi(cx,cy,r,x,y)

	if kondisi == true{
		fmt.Println("Anda berada di dalam taman")
	} else {
		fmt.Println("Anda berada di luar taman")
	}
}

func jarak(a,b,c,d float64) float64{
	return math.Sqrt(((a-c)*(a-c))+((b-d)*(b-d)))
}

func posisi(cx,cy,r,x,y float64) bool{
	var jrk float64

	jrk = jarak(cx,cy,x,y)
	return jrk <= r
}