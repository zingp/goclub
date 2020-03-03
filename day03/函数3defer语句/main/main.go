package main

import (
	"fmt"
)

/*
a. 当函数返回时，执行defer语句。可以用来做资源清理；
b. 多个defer语句，按照先进后出的方式进行；
c. defer 语句中的变量在defer声明时就决定了。
用途：
	关闭文件句柄:defer file.Close()
	锁资源释放:defer mc.Unlock()
	数据库连接释放：defer conn.Close()
*/

func test(a, b int) int {
	//匿名函数的调用
	result := func(a1 int, b1 int) int{
		return a1 + b1
	}(a, b)
	return result
}

func main() {
	var a int = 0
	defer fmt.Println(a)
	defer fmt.Println("lyy")

	a = 10
	fmt.Println(a)
	fmt.Println(test(1,2))
}