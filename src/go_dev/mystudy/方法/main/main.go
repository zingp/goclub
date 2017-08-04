/*
Go 没有类。然而，仍然可以在结构体类型上定义方法。
方法接收者 出现在 func 关键字和方法名之间的参数中。
*/
package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X,Y float64
}

func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type Person struct {
	Name, Country string
	Age int
}

func (p Person) tell() {
	fmt.Printf("My name is %s, i am from %s, and %d years old.\n",p.Name,p.Country,p.Age)
}

func main() {
	v := &Vertex{3,4}
	fmt.Println(v.Abs())

	p := Person{"lyy", "zh_CN", 25}
	p.tell()
}
