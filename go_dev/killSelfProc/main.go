package main


import (
	"time"
	"fmt"
	"sync"
	"os"
	"path"
)

var timeStart = time.Now().Unix()
var waitGroup sync.WaitGroup

// 超时或者日志上传完毕关闭管道，退出各个goroutine
func goroutineExit() {

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {

		if int(time.Now().Unix()-timeStart) > (1 * 60) {
			fmt.Printf("timeout:%d mins\n", 1)
			os.Exit(0)  // 退出进程
			return
		}
	}
}

func main() {
	go func(){
		waitGroup.Add(1)
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("func 1")
		}
	}()
	fmt.Println("after print")
	
	waitGroup.Add(1)
	goroutineExit()
	waitGroup.Wait()
	
}


