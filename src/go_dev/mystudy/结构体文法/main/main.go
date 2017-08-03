package main
/*
结构体文法
*/
import (
	"fmt"
)

type Point struct {
	X, Y int
}

var (
	v1 = Point{1,2}
	v2 = Point{X: 1}
	v3 = Point{}
	p = &Point{1,2}

)

func main() {
	fmt.Println(v1, v2, p, v3)

	p.X = 4
	fmt.Println(p)
}
