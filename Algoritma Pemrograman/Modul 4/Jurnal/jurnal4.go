package main
 
import (
    "fmt"
    "math/rand"
)
 
func main() {
    var seed int64
    var n int
    var tebak1, tebak2 string
 
    fmt.Print("Angka Rahasia: ")
    fmt.Scanln(&seed)
    rand.Seed(seed)
    fmt.Print("Anda: ")
    fmt.Scanln(&tebak1, &tebak2)
    n = rand.Intn(6) + 1
    if n%2 == 1 {
        if n >= 4 {
            if tebak1 == "ganjil" {
                if tebak2 == "kecil" {
                    fmt.Println("Nilai Dadu", n, "Anda kalah")
                } else if tebak2 == "besar" {
                    fmt.Println("Nilai Dadu", n, "Anda menang")
                }
            } else if tebak1 == "genap" {
                fmt.Println("Nilai Dadu", n, "Anda kalah")
            }
        } else if n < 4 {
            if tebak1 == "ganjil" {
                if tebak2 == "kecil" {
                    fmt.Println("Nilai Dadu", n, "Anda menang")
                } else if tebak2 == "besar" {
                    fmt.Println("Nilai Dadu", n, "Anda kalah")
                }
            } else if tebak1 == "genap" {
                fmt.Println("Nilai Dadu", n, "Anda kalah")
            }
        }
    } else if n%2 == 0 {
        if n >= 4 {
            if tebak1 == "ganjil" {
                fmt.Println("Nilai Dadu", n, "Anda kalah")
            } else if tebak1 == "genap" {
                if tebak2 == "kecil" {
                    fmt.Println("Nilai Dadu", n, "Anda kalah")
                } else if tebak2 == "besar" {
                    fmt.Println("Nilai Dadu", n, "Anda menang")
                }
            }
        } else if n < 4 {
            if tebak1 == "ganjil" {
                fmt.Println("Nilai Dadu", n, "Anda kalah")
            } else if tebak1 == "genap" {
                if tebak2 == "kecil" {
                    fmt.Println("Nilai Dadu", n, "Anda menang")
                } else if tebak2 == "besar" {
                    fmt.Println("Nilai Dadu", n, "Anda kalah")
                }
            }
        }
    }
}