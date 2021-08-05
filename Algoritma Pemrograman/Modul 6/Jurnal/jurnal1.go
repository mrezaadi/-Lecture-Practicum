package main

import "fmt"

var per, com int

func main(){

	var n, r int
	
	fmt.Scanln(&n, &r)

	fmt.Println(permutation(n,r), combination(n,r))

}

func findFactorial(x int ) int{
	if x <= 1 {
		return 1
	} else {
		return x * findFactorial(x-1)
	}
}

func permutation(n, r int) int{
	per = findFactorial(n) / findFactorial(n-r)
	return per
}

func combination(n, r int) int{
	com = findFactorial(n) / ((findFactorial(r)) * (findFactorial(n - r)))
	return com

}