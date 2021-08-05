package main

import "fmt"

func main(){
	var plat string
	var bdg, jkt, lain int

	for plat != "."{ //Saat plat bukan "."(mark) masuk ke looping
		fmt.Scan(&plat) //Input plat
		if plat != "."{ //Saat plat bukan ".", lakukan pemanggilan prosedur hitungNopol
			hitungNopol(plat, &bdg, &jkt, &lain)
		}
	}

	fmt.Println(bdg, jkt, lain)
}

func hitungNopol(plat string, bandung, jakarta, lainnya *int){
	if plat == "B" || plat == "F"{ //Jika plat yang di inputkan "B" dan "F" maka termasuk ke daerah Jakarta dan sekitarnya, sehingga jumlah plat Jakarta bertambah 1
		*jakarta++
	} else if plat == "D" || plat == "Z"{ //Jika plat yang di inputkan "D" dan "Z" maka termasuk ke daerah Bandung Raya, sehingga jumlah plat Bandung bertambah 1
		*bandung++
	} else{ //Jika selain "B", "D", "F", dan "Z" maka plat tersebut termasuk ke daerah lainnya, sehingga jumlah plat lainnya bertambah 1
		*lainnya++
	}
}