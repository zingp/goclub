package main

import (
	"os"
	"fmt"
)

//
func echo()  {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	echo()
}