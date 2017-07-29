package main

import "fmt"

/*
数字类型主要有：int/int8/int16/int32/int64/uint8/uint16/uint32/uint64/float32/float64
类型转换，type(var) 比如：var a int8 = 2; var b int32 = int32(a)
*/

func main() {
	var a int8 = 127
	var b int16 = int16(a)
	fmt.Printf("a=%d b=%d", a, b)
}
