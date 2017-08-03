package main

import "fmt"

type Point struct {
	X, Y float64
}
var m map[string]Point

func main() {
	m = make(map[string]Point)
	m["first"] = Point{
		1.00001, 2.00002,
	}
	fmt.Println(m["first"])
}