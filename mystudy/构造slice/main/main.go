package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	printSlice("a", a)
	b := make([]int, 0, 5)
	fmt.Println(b)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)

}
