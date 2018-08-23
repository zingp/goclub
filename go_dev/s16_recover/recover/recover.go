package main

import(
	"fmt"
)

func set(p *int){
	*p = 100
}

func testRecover(){
	// recover 捕获异常，防止程序挂掉；defer 要写在前面
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("Panic:%v\n", err)
		}
	}()

	// 空指针
	var p *int
	set(p)
}

func main(){
	testRecover()
}