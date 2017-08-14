package main
/*用go build -race ... 编译后执行可以查看数据是否有竞争*/
import (
	"math/rand"
	"fmt"
	"time"
	"sync"
)

var lock sync.Mutex  //互斥锁

func testGo() {
	m := make(map[int]int)
	m[1] = 10
	m[3] = 10
	m[8] = 10
	m[11] = 10
    //并发修改map的同一个值
	for i:=0; i<2; i++ {
		rand.Seed(time.Now().UnixNano())
		go func(a map[int]int){
			lock.Lock()
			a[3] = rand.Intn(100)
			lock.Unlock()
		}(m)
	}

	lock.Lock()
	fmt.Println(m)
	lock.Unlock()

	time.Sleep(time.Second)
}

func main() {
	testGo()
}
