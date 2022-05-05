package main

import "fmt"

type Number interface {
	int64 | float64
}


func SumNuber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}


func main() {
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}
	
	// Initialize a map for the float values
	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("ints sum=%d, floats sum=%f\n", SumNuber(ints), SumNuber(floats))
}