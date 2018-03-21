package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//如果两个或者多个 goroutine 在没有互相同步的情况下，访问某个共享的资源，并试图同时
//读和写这个资源，就处于相互竞争的状态，这种情况被称作竞争状态（race candition）。竞争状态
//的存在是让并发程序变得复杂的地方，十分容易引起潜在问题。对一个共享资源的读和写操作必
//须是原子化的，换句话说，同一时刻只能有一个 goroutine 对共享资源进行读和写操作。

var (
	// counter 是所有 goroutine 都要增加其值的变量
	num int64

	// wg 用来等待程序结束
	wg sync.WaitGroup
)

// main 是所有 Go 程序的入口
func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	//go incCounter(1)
	//go incCounter(2)
	go incNum(1)
	go incNum(2)

	// 等待 goroutine 结束
	wg.Wait()
	fmt.Println("Final Num:", num)
}

// incCounter 增加包里 num 变量的值
func incCounter(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 捕获 num 的值
		value := num

		// 当前 goroutine 从线程退出，并放回到队列
		// 用于将 goroutine 从当前线程退出，
		//给其他 goroutine 运行的机会。在两次操作中间这样做的目的是强制调度器切换两个 goroutine，
		//以便让竞争状态的效果变得更明显。
		runtime.Gosched()

		// 增加本地 value 变量的值
		value++

		// 将该值保存回 counter
		num = value
	}
}

//每个 goroutine 都会覆盖另一个 goroutine 的工作。这种覆盖发生在 goroutine 切换的时候。每
//个 goroutine 创造了一个 num 变量的副本，之后就切换到另一个 goroutine。当这个 goroutine
//再次运行的时候， num 变量的值已经改变了，但是 goroutine 并没有更新自己的那个副本的
//值，而是继续使用这个副本的值，用这个值递增，并存回 num 变量，结果覆盖了另一个
//goroutine 完成的工作。


// 用原子函数锁住共享资源
// 原子函数能够以很底层的加锁机制来同步访问整型变量和指针。
func incNum(id int)  {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全地对num 加 1
		atomic.AddInt64(&num, 1)
		runtime.Gosched()
	}
}