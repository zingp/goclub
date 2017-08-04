package main

import (
	"sort"
	"fmt"
)

/*
内置排序
*/

func testIntSort(){
	var a = [...]int{1,34,25,8,300}
	sort.Ints(a[:])      //应该传切片，不应传数组，你懂得
	fmt.Println(a)
}
func testStringSort(){
	var a = [...]string{"abc","ef","A","ghk"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloatSort() {
	var a= [...]float64{2.2, 1.1, 3.5, 0.9}
	sort.Float64s(a[:])
	fmt.Println(a)

	index := sort.SearchFloat64s(a[:],2.2)   //查找
	fmt.Println(index)
}


func main() {
	testIntSort()
	testStringSort()
	testFloatSort()
}