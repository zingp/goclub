package main

import "fmt"

var name string

func func1(){
	name := "zing-p"
	fmt.Println(name)
	func2()
}
func func2(){
	fmt.Println(name)
}

func main(){
	name = "lyy"
	fmt.Println(name)
	//print("123",name)  //最后打印，且写入标准错误？
	func1()
}

/*
result:
lyy
zing-p
lyy
*/