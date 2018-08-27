package main

import (
	"fmt"
)

var Str []string = []string{"ac", "bc", "cc"}

func main() {
	value, ok := interface{}(Str).([]string)
	if !ok {
		fmt.Println("type transfer to []string failed.")
		return
	}

	fmt.Printf("value=%v\n", value)
}
