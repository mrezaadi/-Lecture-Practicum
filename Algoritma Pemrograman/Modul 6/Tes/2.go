package main

import "fmt"

func main(){
	var x int
	var a, b, c string
	var hasil int

	fmt.Scanln(&x, &a, &b, &c)

	if a == "f"{
		if b == "g"{
			if c == "h"{
				hasil = h(g(f(x))) //(hogof)
			}
		} else if b == "h"{
			if c == "g"{
				hasil = g(h(f(x)))//(gohof)
			}
		}
	} else if a == "g"{
		if b == "f"{
			if c == "h"{
				hasil = h(f(g(x))) //(hofog)
			}
		} else if b == "h"{
			if c == "f"{
				hasil = f(h(g(x))) //(fohog)
			}
	} else if a == "h"{
		if b == "f"{
			if c == "g"{
				hasil = g(f(h(x))) //(gofoh)
			}
		} else if b == "g"{
			if c == "f"{
				hasil = f(g(h(x))) //(fogoh)
			}
		}
	}

}
}

func f (x int) int{
	var hasil int

	hasil = (2*x) + 5
	return hasil
}

func g (x int) int{
	var hasil int

	hasil = (x*x) + (2*x)
	return hasil
}

func h(x int) int{
	var hasil int

	hasil = x-3
	return hasil
}
