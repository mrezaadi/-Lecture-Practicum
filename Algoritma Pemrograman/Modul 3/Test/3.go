package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5 int
	var s1, s2, s3, s4, s5 int

	fmt.Scanf("%d%c %d%c %d%c %d%c %d%c", &a1, &s1, &a2, &s2, &a3, &s3, &a4, &s4, &a5, &s5)
	quad := (a2 == a3 && a2 == a4 && a2 == a5)  || (a1 == a3 && a1 == a4 && a1 == a5) || (a4 == a1 && a4 == a2 && a4 == a5) || (a3 == a1 && a3 == a2 && a3 == a5) || (a4 == a1 && a4 == a2 && a4 == a3)

	fmt.Println(quad)
}