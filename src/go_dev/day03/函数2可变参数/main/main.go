package main
/*
注意其中arg是个slice，用len(arg)获取其长度
*/

import "fmt"

func add(a int, arg...int)(sum int){
	sum += a
	for i:=0;i<len(arg);i++ {
		sum += arg[i]
	}
	return
}
func main() {
	res := add(1, 10, 100)
	fmt.Println(res)
}