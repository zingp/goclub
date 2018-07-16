package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var waitGroup sync.WaitGroup
var count int


// 不加锁
func testMutex1() {
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100000; i++ {
			count++
		}
		waitGroup.Done()
	}()

	for i := 0; i < 100000; i++ {
		count++
	}

	waitGroup.Wait()
}


// 加互斥锁 结果正确
func testMutex2() {
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100000; i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
		waitGroup.Done()
	}()

	for i := 0; i < 100000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}

	waitGroup.Wait()
}


func main() {
	testMutex1()
	fmt.Println("count::", count)

}
