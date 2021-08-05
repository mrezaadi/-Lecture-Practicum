package main

import "fmt"

func main(){
	var m, n int 
	var kpk int

	fmt.Scanln(&m, &n) //Input m dan n

	
    if (m>n){ //Jika m > n, kelipatan dimulai dari m
        kpk = m
    } else{ //Jika m < n, kelipatan dimulai dari m
        kpk = n
	}

	for { //Perulangan tanpa kondisi    
        if (kpk%m==0 && kpk%n==0) { //Jika kelipatan habis dibagi oleh m dan n, maka kelipatan tersebut merupakan kpk      
            break //Perulangan berhenti
        }
        kpk++ //Jika kelipatan tidak habis dibagi oleh m dan n, maka kelipatan tersebut bukan kpk, dan kelipayan bertambah 1
	}
	fmt.Println(kpk)
}