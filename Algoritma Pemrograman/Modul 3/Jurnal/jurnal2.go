package main

import "fmt"

func main(){
    var a,b,c,d int
    
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Scanln(&c)
    fmt.Scanln(&d)

    fmt.Println(a == d)
    fmt.Println(b == (a/1000) || b == (a/100)%10 || b == (a/10)%10 || b == (a%10))
    fmt.Println(c == (a/100) || c == (a/10)%100 || c == (a%100))
}