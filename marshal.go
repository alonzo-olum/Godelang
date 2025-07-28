package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
	h := json.RawMessage(`{"processed": true}`)

	c := struct {
            Header *json.RawMessage	`json: "header"`
	    Body string			`json: "body"`
    }{Header: &h, Body: "Document Complete!"}

    b, err := json.MarshalIndent(&c, "", "\t")
    if err != nil {
	    fmt.Println("errors:", err)
    }
    os.Stdout.Write(b)
}

