package main

import (
	"time"
	"sync"
	"math/rand"
	"fmt"
)

/*
多读写少的情况下，用读写锁。
这种情况读写锁的性能比互斥锁高约100倍
*/
var rwLock sync.RWMutex

func testRWLock(){
	m := make(map[int]int,100)
	m[1] = 10
	m[3] = 10
	m[8] = 10
	m[11] = 10
	//并发修改map的同一个值
	for i:=0; i<2; i++ {
		rand.Seed(time.Now().UnixNano())
		go func(a map[int]int){
			rwLock.Lock()
			a[3] = rand.Intn(100)
			rwLock.Unlock()
		}(m)
	}

	for i:=1;i<100;i++ {
		go func(a map[int]int) {
			rwLock.RLock()
			fmt.Println(a)
			rwLock.RUnlock()
		}(m)
	}
	time.Sleep(time.Second*3)
}

func main() {
	testRWLock()
}