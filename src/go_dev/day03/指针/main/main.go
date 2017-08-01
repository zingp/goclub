package main

import "fmt"
/*
搞清楚指针变量--地址--值之间的关系。
*/
func main() {
	var a int = 10
	println(&a)

	var p *int
	p = &a

	fmt.Println(*p)
	*p = 100
	fmt.Println(a)
}
