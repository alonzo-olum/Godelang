package main

import "fmt"

func append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
	}
	slice = slice[0:len(data)]
	copy(slice[1:], data)
	return slice
}

func main() {
	s := make([]byte, 10)
	d := []byte{'a','b','c'}
	s = append(s, d)
	fmt.Printf("%c\n", s)
}
