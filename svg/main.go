package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var handler = func(w http.ResponseWriter, r *http.Request){
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	w.Header().Set("Content-Type", "image/svg+xml")
	// initialize sentinels
	color := "blue"
	width, height := 600, 320
	//extract query params
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(params["color"]) > 0 {
		color = params["color"][0]
		fmt.Printf("?color=%s\n", color)
	}
	svg(w, color, int(height), int(width))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
