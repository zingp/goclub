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


//头插法:要改变指针变量的值，需要传指针变量的指针进去
func insertHead(p **Student) {
     for i:=0; i<10; i++ {
		 stu := Student{
			 Name: fmt.Sprintf("stu%d", i),
			 Age: rand.Intn(100),
		 }
		 stu.next = *p
		 *p = &stu
	 }
}

func insertHeadTest() {
	var head *Student = new(Student)
	head.Name = "lyy"
	head.Age = 25

	insertHead(&head)
	traversal(head)
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

	insertHeadTest()
}
