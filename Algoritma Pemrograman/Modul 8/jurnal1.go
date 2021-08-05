package main

import "fmt"

func main(){
	var n,p,q int

	fmt.Scanln(&n, &p, &q) //Input nilai n,p,dan q

	if p == q{ //Saat kondisi p dan q sama
		fmt.Println("Kita tidak pernah bertemu")
	} else if (n%p == 0 || n%q == 0) && !(n%p == 0 && n%q == 0){ //Saat salah satu p atau q habis membagi n, tetapi tidak keduanya (XOR)
		fmt.Println("Kita bertemu hari ini")
	} else if n%p == 0 && n%q == 0 { //Saat p dan q keduanya habis membagi n
		if p<q{ //Jika p < q, maka yang digunakan p
			fmt.Println("Kita akan bertemu", p, "hari lagi")
		} else{ //Jika p > q, maka yang digunakan q
			fmt.Println("Kita akan bertemu", q, "hari lagi")
		}
	} else{ //Saat p dan q keduanya tidak habis membagi n
		a := p - n%p
		b := q - n%q
		if a<b{ //Saat a < b, maka yang digunakan a
			fmt.Println("Kita bertemu",a,"lagi")
		}  else{ //Saat b < a, maka yang digunakan b
			fmt.Println("Kita bertemu",b,"lagi")
		}
	}
}