package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex
var waitGroup sync.WaitGroup
var count int
/*读多写少用读写锁*/

func testLock(){
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			lock.Lock()
			count++
			time.Sleep(5*time.Millisecond)
			lock.Unlock()
		}
		waitGroup.Done()
	}()

	for i:=0;i<16;i++{
		waitGroup.Add(1)
		go func(){
			for i := 0; i < 5000; i++ {
				lock.Lock()
				// 模拟读  耗时1ms
				time.Sleep(1*time.Millisecond)
				lock.Unlock()
			}
			waitGroup.Done()
		}()	
	}

	waitGroup.Wait()
}

func testRwLock(){
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			rwLock.Lock()
			count++
			time.Sleep(5*time.Millisecond)
			rwLock.Unlock()
		}
		waitGroup.Done()
	}()

	for i:=0;i<16;i++{
		waitGroup.Add(1)
		go func(){
			for i := 0; i < 5000; i++ {
				// 读锁
				rwLock.RLock()
				// 模拟读  耗时1ms
				time.Sleep(1*time.Millisecond)
				rwLock.RUnlock()
			}
			waitGroup.Done()
		}()	
	}

	waitGroup.Wait()
}

func main() {
	start := time.Now().UnixNano()
	testLock()
	// testRwLock()
	end := time.Now().UnixNano()
	fmt.Println("cost time:", (end-start)/1000/1000)
}