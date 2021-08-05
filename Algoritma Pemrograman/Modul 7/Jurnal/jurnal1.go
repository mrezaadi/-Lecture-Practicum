package main

import "fmt"

func main() {
    var biner string
    var desimal int

    fmt.Scanln(&desimal)
    biner = Des2Bin(desimal)
    fmt.Println(biner)
}

func Division(a, b int, result, remainder *int) {
    *result = a / b
    *remainder = a % b
}

func Num2Str(x int) string {
    if x == 0 {
        return "0"
    } else if x == 1 {
        return "1"
    } else {
        return ""
    }
}

func Des2Bin(desimal int) string {
    var rDiv, rMod int
    hasil := ""
	if desimal == 0{ //Disini saya sedikit menambahkan perubahan, karena saat saya gunakan pseudocode, saat input 0 tidak sesuai dengan ouput contoh
		return "0"
	} else{
		for desimal > 0 {
			Division(desimal, 2, &rDiv, &rMod)
			desimal = rDiv
			hasil = Num2Str(rMod) + hasil
		}
		return hasil
	}

}