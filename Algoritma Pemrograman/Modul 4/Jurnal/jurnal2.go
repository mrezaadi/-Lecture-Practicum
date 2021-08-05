package main

import "fmt"

func main(){
	var n int
	var nilai int
	var max, jmlh_max int

	max = 0

	fmt.Scanln(&n)
	for i := 1; i <= n; i++{
		fmt.Scan(&nilai)
		if nilai > max {
			max = nilai
			jmlh_max = 1
		} else if nilai == max{
			jmlh_max++
		}
	}
	fmt.Println("Nilai terbesar adalah", max, "yang diperoleh oleh", jmlh_max, "orang mahasiswa")
}