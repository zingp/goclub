package main

import (
	"../testPipe"
	"fmt"
)
/*
导入包才能引用包里的公有成员，公有成员的第一个字母是大写；
从管道中取出值<- pipe。
*/
func main(){
	pipe := make(chan int, 3)
	go testPipe.AddFunc(100, 200, pipe)
	res1 :=<- pipe
	res2 :=<- pipe
	res3 :=<- pipe
	fmt.Println(res1, res2, res3)

	// 如果初始化一个有缓冲管道，不放数据直接读会自杀
	pipe2 := make(chan int, 10)
	pipe2 <- 8
	r :=<- pipe2
	fmt.Println(r)
}
