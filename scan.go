package main

import (
	"fmt"
 	"unicode"
)

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])); i++ {
	}
	x := 0
	for ; i < len(b) && unicode.IsDigit(rune(b[i])); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func main() {
	b := []byte{'1', '2', 'q'}
	for i := 0; i < len(b); {
		x, _ := nextInt(b, i)
		fmt.Println(x)
	}
}
