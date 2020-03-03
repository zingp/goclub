package main

import (
	"fmt"
	"time"
	"sync"
)

var testChan = make(chan int, 10)
var stopChan = make(chan int, 1)     //设置一个管道，控制携程退出
var waitGroup sync.WaitGroup


func producer() {
	defer waitGroup.Done()
	defer fmt.Println("producer ok")

	for i:=0;i<100;i++ {
		time.Sleep(time.Second)

		select {
		case _, ok :=<- stopChan:
			if !ok {
				fmt.Printf("stop !ok: %v\n", ok)
				return
			}
		case testChan <- i:
		default:
			fmt.Printf("defaut: %v\n", i)
		}

	}
}

func producer2() {
	defer waitGroup.Done()
	defer fmt.Println("producer ok")

	for i:=0;i<100;i++ {
		time.Sleep(time.Second)

		select {
		case <- stopChan:
			fmt.Println("producer2 recieve a top signal!")
			return
		case testChan <- i*10:
		default:
			fmt.Printf("defaut: %v\n", i)
		}
		
	}
}

func customer() {
	defer waitGroup.Done()

	for {
		select {
		case res, ok :=<- testChan:
			if !ok {
				fmt.Printf("costom !ok:, %v\n", ok)
			}
			fmt.Printf("costom res:, %d\n", res)
		case  <- stopChan:
			fmt.Println("costom recieve a stop signal!")
			return
		}
	}
}

func exit(){
	defer waitGroup.Done()
	defer fmt.Println("exit....")

	time.Sleep(time.Second * 10)
	close(stopChan)
}

func main(){
	waitGroup.Add(1)
	go producer()

	waitGroup.Add(1)
	go producer2()

	waitGroup.Add(1)
	go customer()

	waitGroup.Add(1)
	go exit()

	waitGroup.Wait()
}

// https://studygolang.com/articles/16646