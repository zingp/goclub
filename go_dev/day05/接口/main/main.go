package main

import "fmt"
//interface类型默认是一个指针
type Test interface {
	Tell()
	Sleep()
}

type People struct {
	Name string
	Age int
}

type Student struct {
	People
	Score int
}

func (p *People) Tell() {
	fmt.Printf("My name is %s.\n", p.Name)
}

func (p *People) Sleep()  {
	fmt.Println("People am sleeping..")
}

func (p *Student) Tell() {
	fmt.Printf("My score is %d.\n", p.Score)
}

func (p *Student) Sleep()  {
	fmt.Println("Student is sleeping.")

}

func main() {
	var t Test
	var p01 People
	var stu Student
	p01.Name = "zing-p"
	p01.Age = 25

	stu.Name = "lyy"
	stu.Age = 25
	stu.Score = 99

	t = &p01
	t.Tell()
	t.Sleep()

	t = &stu
	t.Tell()
	t.Sleep()
}


