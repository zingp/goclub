package main

import(
	"fmt"
	"sync"
)

func testChan(){
	// 定义： var 变量名 chan  类型
	var intChan chan int

	// 初始化  如果不初始化就使用报错
	intChan = make(chan int, 10)  // 有缓冲区。如果无缓冲区无人来取时会阻塞
	
	go func(){
		intChan <- 9
	}()

	result :=<- intChan
	fmt.Printf("result=%d\n", result)
}


func producer(s chan string){
	for i:=0;i<3;i++ {
		s <- fmt.Sprintf("item%d", i)
	}
	close(s)
	waitGroup.Done()
}

func comsumer(s chan string) {
	for {
		str, ok :=<- s
		if !ok {
			fmt.Printf("chan is closed")
			break
		}
		fmt.Printf("value=%s\n", str)
	}
	waitGroup.Done()
}

var waitGroup sync.WaitGroup
func main() {
	// testChan()

	var strChan chan string = make(chan string, 10)
	waitGroup.Add(2)
	go producer(strChan)
	go comsumer(strChan)
	waitGroup.Wait()
}