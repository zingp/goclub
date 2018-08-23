package main

import "fmt"
// make用来创建map/slice/channel
// new用来创建值类型

func test() {
	s1 := new([]int)  //slice,返回指针
	fmt.Println(s1)

	s2 := make([]int, 10)
	fmt.Println(s2)

	*s1 =make([]int,5)  //初始化
	(*s1)[0] = 100
	s2[0] = 2

	fmt.Println("s1=",s1)
	fmt.Println("s2=",s2)
}

func main() {
	test()
}
