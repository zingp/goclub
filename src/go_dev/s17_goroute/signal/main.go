package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var waitGroup sync.WaitGroup

func poducer(strChan chan string, flagChan chan bool) {
	var (
		i        int
		exitFlag bool
	)
	// 生产者不断往队列中push数据
	for {
		str := fmt.Sprintf("chan data %d", i)
		select {
		case strChan <- str:
		case exitFlag = <-flagChan:
		}
		if exitFlag {
			fmt.Printf("exitFlag=%v", exitFlag)
			break
		}
	}

	close(strChan)
	waitGroup.Done()
}

func consumer(strChan chan string) {
	for {
		str, ok := <-strChan
		if !ok {
			fmt.Println("chan is closed.\n")
			break
		}
		fmt.Printf("get data for chan:%s", str)
	}

	waitGroup.Done()
}


// 使用方法：在Linux 下先ps出进程号，然后kill SIGUSR2  进程ID
func main() {
	var strChan chan string = make(chan string)
	var flagChan chan bool = make(chan bool, 1)
	var sigChan chan os.Signal = make(chan os.Signal, 1)

	waitGroup.Add(2)

	// syscall.SIGUSR2 在linux下可用
	signal.Notify(sigChan, syscall.SIGUSR2)  // 注册信号，给进程发送信号

	go poducer(strChan, flagChan)
	go consumer(strChan)

	<-sigChan
	flagChan <- true

	waitGroup.Wait()

}
