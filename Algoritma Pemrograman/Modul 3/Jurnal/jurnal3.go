package main

import "fmt"

func main() {
    var a1, a2, a3, a4, a5 int
    var s1, s2, s3, s4, s5 string

    fmt.Scanln(&a1, &s1, &a2, &s2, &a3, &s3, &a4, &s4, &a5, &s5)

    straight := (a2-a1 == 1 && a1 != 1) && (a3-a2 == 1) && (a4-a3 == 1) && (a5-a4 == 1 || a5 == 1)
    flush := (s1 == s2) && (s1 == s3) && (s1 == s4) && (s1 == s5)

    fmt.Println(straight)
    fmt.Println(flush)

}