package main

import (
	"fmt"
	"time"
)

/*
一个函数调用自己就叫递归
问题可以拆分成很多个子问题
问题的规模在不断缩小
递归有明确的结束条件（出口条件）
    go的堆栈很廉价
*/
func recursive(n int){
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n >10 {
		return
	}
	recursive(n+1)
}

//计算阶乘？
func calc(n int) int {
	if n == 1{
		return 1
	}
	return calc(n-1)*n
}

//斐波那契数列的通项
func fab(n int) int {
	if n <= 1{
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {

	for i:= 1;i<10;i++ {
		fmt.Println(fab(i))   //求fab数列
	}

	res := calc(5)
	fmt.Printf("5的阶乘是%d\n", res)

	//recursive(5)
}