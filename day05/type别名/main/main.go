package main

import "fmt"

//自定义类型，但不等同int
type integer int

func main() {
	var a integer = 1000
	var b int
	b = int(a)  // 强制转换一下。

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
}
