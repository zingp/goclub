package main

import "fmt"
/*
for i, v := range 语法用来遍历数组/slice/map/chan
*/
func main() {
	str := "hello world, 中国"
	for i, v := range str{
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(string(v))))
	}
}