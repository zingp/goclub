package main

import (
	"runtime"
	"sync"
	"fmt"
)


// 用来等待程序完成
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create 2 goroutine \n")
	go printPrime("A")
	go printPrime("B")

	// 等待goroutine结束
	fmt.Println("Waiting finish\n")
	wg.Wait()

	fmt.Println("Terminating Programe")
}

// 输出5000以内的素数
func printPrime(str string) {
	defer wg.Done()

next:
	for outer := 2;outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", str, outer)
	}
	fmt.Println("Completed", str)
}
