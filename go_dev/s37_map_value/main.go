package main

import (
	"fmt"
)

/*
测试修改map的值
map 是引用传递
*/
var m = make(map[string]int)
func initMap() {
	m["age"] = 23
	m["score"] =100
}
func modify(x map[string]int){
	x["value"] = 1
}

func main() {
	fmt.Println(m)
	modify(m)
	fmt.Println(m)
}