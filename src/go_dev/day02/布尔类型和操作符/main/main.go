package main

import "fmt"

/*
bool类型只能存false和true,默认false
相关操作符：!, &&, ||
*/

func main() {
	var a bool
	var b = true
	fmt.Println(a)
	fmt.Println(!a, !b, a && b, a || b)
}
/*
false
true false false true
*/