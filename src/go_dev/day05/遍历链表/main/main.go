package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age int
	next *Student
}

// 链表遍历函数
func traversal(p *Student) {
	for p != nil {
		fmt.Println(p.Name)
		p = p.next
	}
}

//尾插法
func insertTail(p *Student) {
	var tail = p
	for i:=0;i<10;i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age: rand.Intn(100),
		}
		tail.next = &stu
		tail = &stu
	}
}

func main() {
	var stu1 = Student{
		Name:"lyy",
		Age:25,
	}
	var stu2 = Student{
		Name:"zing-p",
		Age:25,
	}
	stu1.next = &stu2

	insertTail(&stu2)
	traversal(&stu1)
}
