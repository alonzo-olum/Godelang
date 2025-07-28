package main

import "fmt"
import "os"
import "bufio"

func counter(input *bufio.Scanner) {
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	counter(input)
}
