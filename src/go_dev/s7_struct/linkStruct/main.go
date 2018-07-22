package main

import (
	"fmt"
)

// go链式编程  返回对象
type Person struct {
	Name string
	Sex string
	Age int
}

func (p *Person)SetName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person)SetSex(sex string) *Person {
	p.Sex = sex
	return p
}

func (p *Person)SetAge(age int) *Person {
	p.Age = age
	return p
}

func (p *Person) Print() {
	fmt.Printf("Name is %s\nSex is %s\nAge is %d\n",p.Name,p.Sex,p.Age)
}

func main(){
	p := &Person{}
	p.SetName("zingp").SetSex("male").SetAge(25).Print()
}