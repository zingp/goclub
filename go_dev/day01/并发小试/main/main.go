package main
/*
本模块主要体验go语言的并发。
通过go这个关键字就可以实现。
注意要让go的主进程等待goroutine结束。
*/
import (
	"fmt"
	"time"
)

func main() {
	for i:=1; i<100; i++ {
		go fmt.Println(i)
	}
	// 让主进程等待gorotine结束
	time.Sleep(time.Second)
}
