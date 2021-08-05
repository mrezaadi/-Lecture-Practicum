package main

import "fmt"

func main() {
    var k, n1, n2 int
    fmt.Scanln(&k)

    for n1 != -1 && n2 != -1 {
        fmt.Scanln(&n1, &n2)
        if n1%k != 0 {
            fmt.Println(n1)
        } else if n2%k != 0 {
            fmt.Println(n2)
        }
    }
}