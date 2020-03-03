package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


var pChan = make(chan int, 10)
var cChan = make(chan int, 10) 
var waitGroup sync.WaitGroup

var flagNum int32

func main(){

	// waitGroup.Add(8)
	// for i:=0;i<4;i++ {
	// 	go producer()
	// }

	// for i:=0;i<4;i++ {
	// 	go consumer()
	// }
	// waitGroup.Wait()
	// fmt.Println("done")
	var m = make(map[string]string, 10)
	m["abc"] = "123"

	fmt.Println(len(m))
}

func producer(){
	for i:=0;i<4;i++ {
		atomic.AddInt32(&flagNum, 1)
		n := flagNum
		pChan <- int(n)
	}
	waitGroup.Done()
}

func consumer(){
	for item := range pChan {
		fmt.Println("Consumer:", item)

	}
	waitGroup.Done()
}