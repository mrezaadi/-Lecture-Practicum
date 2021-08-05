package main

import "fmt"

func main(){
	var n float64
	var y float64

	fmt.Scanln(&n)

	fmt.Println("f(", n, ") = ", f(n))
	fmt.Println("g(", n, ") = ", g(n))
	fmt.Println("h(", n, ") = ", h(n))

	komposisi(n, &y)
	fmt.Print(y)
}

func komposisi(x float64, y *float64){
	*y = f(g(h(x)))
}	

func f (x float64) float64{
	var hasil float64

	hasil = (x*x)
	return hasil
}

func g (x float64) float64{
	var hasil float64

	hasil = x-2
	return hasil
}

func h (x float64) float64{
	var hasil float64

	hasil = x + 1
	return hasil
}