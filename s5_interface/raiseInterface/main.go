package main

import (
	"fmt"
)

func typeRaise() {
    var t int = 100
    var x interface{}
    x = t
    y := x.(int)   // 强制转换
    fmt.Println(y)
}

func typeRaiseJustify() {
    var t int = 100
    var x interface{}
    x = t
    y, ok := x.(string)   // 强制转换
    if !ok {
        fmt.Println("convert faild!\n")
    }

    fmt.Println(y)
}

func main() {
    typeRaise()         // 强制转换
    typeRaiseJustify()  //带判断类型转换
}