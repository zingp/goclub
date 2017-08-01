package main

import "fmt"

/*
golang函数的特点：
a.不支持重载：一个包不能有两个名字一样的函数；
b.函数是一等公民，函数也是一种类型，一个函数可以赋值给变量；
c.匿名函数
d.多值返回
*/

type add_func func(int, int) int   //自定义类型

func add(a int, b int)int {
	return a +b
}
func operator(op add_func, a int, b int) int {
	return op(a, b)
}

func calc(a, b int) (sum int, avg int){
	//定义返回值的名字
	sum = a + b
        avg = (a + b)/2
	return
}

func main() {
	c := add
	fmt.Println(c)

	res := operator(c, 100, 200)
	fmt.Println(res)

	_, avg := calc(100, 200)  // 用_接收不使用的返回值
	fmt.Println(avg)
}