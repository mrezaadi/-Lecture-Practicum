package main

import "fmt"

func main(){
	var n, i int
	var mi, mo, ma int
	var jum_ma int
	var tahun int

	fmt.Scanln(&n)
	jum_ma = 0
	i = 1
	tahun = 0

	for i <= n{
		fmt.Scanln(&mi,&mo,&ma)
		if (jum_ma + (mi-mo) != ma) && tahun == 0 {
			tahun = i
		}
		jum_ma = ma
		i = i + 1
	} 
	fmt.Println(tahun)
}