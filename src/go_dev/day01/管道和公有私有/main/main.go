package main

import (
	"go_dev/day01/管道和公有私有/testPipe"
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
}
