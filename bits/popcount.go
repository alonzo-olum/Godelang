package main

import (
	"fmt"
	"strconv"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("pc[%d/2] %d + %v\n", i, pc[i/2], byte(i&1))
	}
}

func popCount(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(uint8(i)*8))]
	}
	return int(b)
}

func popCountShift(x uint64) int {
	pop := 0
	for i := uint(0); i < 64; i ++ {
		if x&(1<<i) != 0 {
			pop++
		}
	}
	return pop
}

func popCountClear(x uint64) int {
	pop := 0
	for x != 0 {
		x = x & (x-1)
		pop++
	}
	return pop
}

func main() {
	x, _ := strconv.ParseUint(os.Args[1], 10, 64)
	fmt.Println(popCount(x))
	fmt.Println(popCountShift(x))
	fmt.Println(popCountClear(x))
}
