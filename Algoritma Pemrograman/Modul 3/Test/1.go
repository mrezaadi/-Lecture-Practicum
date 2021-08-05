package main

import "fmt"

func main() {
    var a, b, c, d, e byte
    var geser int
    fmt.Scanf("%c%c%c%c%c", &a, &b, &c, &d, &e)
    fmt.Scan(&geser)

    a = a + byte(geser)
    b = b + byte(geser)
    c = c + byte(geser)
    d = d + byte(geser)
    e = e + byte(geser)

    fmt.Printf("%c%c%c%c%c", a, b, c, d, e)
}