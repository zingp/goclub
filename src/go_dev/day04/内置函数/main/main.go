package main

import (
	"fmt"
	"time"
)

func test(){
	defer func(){
		if err := recover();err != nil {
			fmt.Println(err) //查一下debug包，把堆栈信息打出来
		}
	}()

	b := 0
    a := 1/b
	fmt.Println(a)
	return
}

func main() {
	var i int
	fmt.Println(i)

	j := new(int)    // 申请值类型的地址（栈分配），得到的是指针
	fmt.Println(j)

	var a [5]int
	fmt.Println(a)

	var b []int
	b = append(b, 10, 20)
	fmt.Println(b)
	b = append(b, b...)
	fmt.Println("b =",b)


	for {
		test()
		time.Sleep(time.Second)
	}

}
