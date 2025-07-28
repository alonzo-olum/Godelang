package main

import "fmt"

func min(a ...int) int {
	min := int(^uint(0) >> 1)
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {
	arr := []int{1,2,3}
	fmt.Println(min(arr...))
}
