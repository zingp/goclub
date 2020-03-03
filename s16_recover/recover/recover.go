package main

import(
	"fmt"
	// "runtime/debug"
)


func myfunc() {
	fmt.Println("a")
    panic(55)
    fmt.Println("b")
	fmt.Println("f")
}

func main(){
	// recover 捕获异常，防止程序挂掉；defer 要写在前面
	defer func(){

		if err := recover(); err != nil {
			fmt.Printf("Panic:%v\n", err)
		} else {
			fmt.Println("process exited normally with status 0.")
		}
	}()

	myfunc()
}