package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer func() {
		resp.Body.Close()
		if e := recover(); e != nil {
			err = nil
			fmt.Printf("we aim to recover!")
		}
	}()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("fetch took: %.2fs %7d %s", secs, data, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	out, err := os.OpenFile("./out", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	for range os.Args[1:] {
		writer.WriteString(<-ch)
		writer.Flush()
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
