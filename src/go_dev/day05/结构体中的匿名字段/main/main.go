package main

import (
	"time"
	"fmt"
)

type Cart struct {
    name string
    age int
}

type train struct {
	Cart
	int
	start time.Time
}

func main() {
	var t train
	t.Cart.name = "train01"    // 可以通过匿名字段的类型去访问
	t.Cart.age = 22
	t.int = 15

	t.name = "001"       // 如果匿名字段类型是结构体，可省略类型
	t.age = 2

	fmt.Println(t)
}

