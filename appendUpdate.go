package main

import "fmt"

type ByteSlice []byte

// Implements the Writer interface
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
	}
	slice = slice[0:len(data)]
	copy(slice[1:], data)
	*p = slice
	return len(data), nil
}

func main() {
	slice := make(ByteSlice, 100)
	/*
	data := []byte{'a','b','c'}
	if n, err := slice.Write(data); err == nil {
		fmt.Printf("ByteSlice: %c len: %d\n", slice, n)
	}
	*/
	fmt.Fprintf(&slice, "This hour has %d days\n", 7)
	fmt.Printf("ByteSlice: %c\n", slice)
}
