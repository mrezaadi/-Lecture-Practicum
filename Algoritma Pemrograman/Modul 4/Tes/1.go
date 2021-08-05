package main 

import "fmt"

func main(){
	    var n int //Input angka
		var i int //Iterasi
		var jmlh_faktor int //Jumlah Faktor
    fmt.Scanln(&n)
    i = 1
    for i <= n {
        if n%i == 0 {
            fmt.Print(i, " ")
            jmlh_faktor = jmlh_faktor + 1
        }
        i++
    }
    if jmlh_faktor == 2 {
        fmt.Println("\nOleOle")
    } else {
        fmt.Println("\nBukan OleOle")
    }
}