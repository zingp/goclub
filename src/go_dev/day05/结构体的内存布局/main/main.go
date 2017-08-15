package main

import "fmt"

type Student struct {
	Name string
	Age int32
	score float32
}

func main() {
	var stu Student
	stu.Name = "zing-p"
	stu.Age = 25
	stu.score = 80.0

	fmt.Println(stu)
	fmt.Printf("%p\n",&stu.Name)    //占10个字节
	fmt.Printf("%p\n",&stu.Age)     //占4个自己
	fmt.Printf("%p\n", &stu.score)
	//另一种初始化方法:
	var stu1 *Student = &Student{
		Name:"lyy",
		Age:25,
	}
	fmt.Println(stu1.Name)

	//另一种初始化方法:
	var stu2 = Student{
		Name:"LiuYY",
		Age:25,
	}
	fmt.Println(stu2.Name)
}