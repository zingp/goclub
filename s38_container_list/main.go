package main

import (
	"fmt"
	"container/list"
)

func main() {
	li := list.New()
	for i:=0;i<100;i++ {
		li.PushBack(i)
	}

	// 遍历
	for e := li.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// fmt.Printf("container/list li:%v", li)
}