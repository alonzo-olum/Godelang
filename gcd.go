package main

import "fmt"

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	x, y := 20, 15
	fmt.Println(gcd(x, y))
}
