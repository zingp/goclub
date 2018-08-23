package main

import "fmt"

/*
数组：同一种数据类型的固定长度的序列
定义： var a [len]int 比如var a [5]int
长度是数组类型的一部分，因此，ar a [5]int；ar a [10]int是不同类型
数组可以通过下标进行访问，下标从0开始，最后一个元素下标是len-1
访问越界：访问下标在数组下标范围外，会触发越界，会panic
数组是值类型，因此改变副本的值，不会改变本身的值。
*/

func main() {
	var a [10]int
	a[0] = 10
	a[1] = 100

	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}

	for index,v := range a {
		fmt.Printf("a[%d] = %d, ", index, v)
	}
}

