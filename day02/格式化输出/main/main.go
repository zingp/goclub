package main

import "fmt"
/*
%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
%#v	相应值的Go语法表示
%T	相应值的类型的Go语法表示
%%	字面上的百分号，并非值的占位符
*/
func main() {
	var a int
	var b bool
	var c byte = 'c'

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
