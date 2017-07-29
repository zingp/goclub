package main
/*
局部变量：函数内部声明的变量叫做局部变量，生命周期仅限于函数内部，
          代码块内部的变量外部也不能访问。
全局变量：函数外部声明的变量叫做全局变量；生命周期作用于整个包，
	  如果首字母大写作用于整个程序。
*/

import "fmt"

var name string  = "lyy"

func func1(){
	name := "zing-p"
	fmt.Println(name)
}

func func2(){
	name = "zing-p"
	fmt.Println(name)
}

func main()  {
	func1()
	fmt.Println(name)
	func2()
	fmt.Println(name)
}

/*
运行结果：
zing-p
lyy
zing-p
zing-p
*/