package main

import "fmt"

type Stringer interface {
	String() string
}

type value string

func (v value) String() string {
	return string(v)
}

func checkStringer(i interface{}) string {
	var str string
	switch value := i.(type) {
	    case string:
		    str = value
	    case Stringer:
		    str = value.String()
	}
	return str
}

func main() {
	var val value = "amen! please work!"
	fmt.Println(checkStringer(val))
	
}

