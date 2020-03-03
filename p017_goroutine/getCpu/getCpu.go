package main

import(
	"fmt"
	"runtime"
	"time"
	"sync"
)


func getCpuNum(){
	n := runtime.NumCPU()
	fmt.Printf("CPU num is %d.", n)

	runtime.GOMAXPROCS(n-1)  // 设置goroute 跑在多少个核上
}

var waitGroup sync.WaitGroup

func do() {
	time.Sleep(time.Second * 2)
	waitGroup.Done()  // 每个goroute执行完，调用waitGroup.Done() 相当于减1
}

func main(){
	// getCpuNum()
	start := time.Now().UnixNano()
	for i:=0; i<5; i++ {
		waitGroup.Add(1)  // 每个goroute执行前，调用waitGroup.Add(1) 相当于加1
		go do()
	}

	waitGroup.Wait()      // 等待所有goroute执行完
	end := time.Now().UnixNano()

	fmt.Printf("cost time:%d\n", (end-start)/1000/1000)
}