package main

import "fmt"

/*
Go 中的方法是作用在特定类型的变量上，因此自定义类型，都可以有方法，而不仅仅是struct
方法的文法：func (receiver type) methodName(参数列表)(返回值列表){}
*/

type Student struct {
	Name string
	Age int
	Score int
}
// 永远记住，如果要改变值，需要传指针
func (p *Student) init(name string, age int, score int) {
	p.Age = age
	p.Name = name
	p.Score = score
}

// 如果不传指针，无法改变值
func (p Student) changeScore(score int) {
	p.Score = score
}

type integer int
// 自定义类型也可以定义方法
func (p *integer) set (i int){
	*p = integer(i)
}

func main() {
	var stu Student
	//(&stu).init("zing-p", 25, 99) 简化为下面的写法
	stu.init("zing-p", 25, 99)
	fmt.Println(stu)

	stu.changeScore(100)
	fmt.Println(stu)

	var a integer = 88
	a.set(99)
	fmt.Printf("a=%d",a)
}