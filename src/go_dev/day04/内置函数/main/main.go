package main

import (
	"fmt"
	"time"
)
/*
close用来关闭channel
len用来求长度，比如string/arrav/silice/map/channel
new 用来分配值类型，比如int，结构体，返回指针，需要初始化
make 用来分配内存，主要用来分配应用类型，比如chan/map/silce
append用来追加元素到数组、slice
panic和recover做错误处理
*/
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
