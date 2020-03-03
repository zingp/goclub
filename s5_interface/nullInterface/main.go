package main

import(
	"fmt"
)

func main() {
	var a interface {}
	var b int
	b = 1000

	a = b
	fmt.Println(a)
}

