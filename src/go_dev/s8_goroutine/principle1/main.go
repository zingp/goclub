package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// 给每个可用的核心分配一个逻辑处理器
	//runtime.GOMAXPROCS(runtime.NumCPU())

	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start a goroutine...")

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 等待goroutine结束
	fmt.Println("Waiting to finish.")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

/*
调用了 runtime 包的 GOMAXPROCS 函数。这个函数允许程序
更改调度器可以使用的逻辑处理器的数量。如果不想在代码里做这个调用，也可以通过修改和这
个函数名字一样的环境变量的值来更改逻辑处理器的数量。给这个函数传入 1，是通知调度器只
能为该程序使用一个逻辑处理器。
*/
