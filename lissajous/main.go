package main

import (
	"log"
	"net/http"
	"fmt"
	"strconv"
	"os"
)

var handler = func(w http.ResponseWriter, r *http.Request){
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	form := make(map[string][]string)
	for k, v := range r.Form {
		form[k] = append([]string{}, v...)
		fmt.Printf("Form[%q] = %q\n", k, v)
	}

	if cycles, ok := form["cycles"]; ok {
		fmt.Printf("cycles %v\n", cycles[0])
		cycle, _ := strconv.Atoi(cycles[0])
		lissajous(w, cycle)
	}
}

func main() {
	// render output to stdout
	lissajous(os.Stdout, 0)
	// render output on the browser
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

