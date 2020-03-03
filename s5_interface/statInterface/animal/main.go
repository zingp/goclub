package main

import(
	"fmt"
)

//接口规范两个方法
type Animal interface{
	Eat() 
	Talk()
}

// 定义Dog结构体，并实现Animal接口中的所有方法
type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating...\n", d.Name)
}

func (d *Dog) Talk() {
	fmt.Printf("%s is talking...\n", d.Name)
}


// 定义Cat结构体，并实现Animal接口中的所有方法
type Cat struct {
	Name string
}

func (c *Cat) Eat() {
	fmt.Printf("%s is eating...\n", c.Name)
}

func (c *Cat) Talk() {
	fmt.Printf("%s is talking...\n", c.Name)
}


func justify() {
	d := &Dog {
		Name:"BUCKER",
	}
	var v interface{} = d
	if dog, ok := v.(Animal); ok {
		dog.Eat()
	}
}


func main() {
	justify()
	// var myAnimal []Animal
	// d1 := &Dog{
	// 	Name:"阿黄",
	// }
	// myAnimal = append(myAnimal, d1)

	// d2 := &Dog{
	// 	Name:"旺财",
	// }
	// myAnimal = append(myAnimal, d2)

	// c1 := &Cat{
	// 	Name:"小七",
	// }
	// myAnimal = append(myAnimal, c1)

	// for _, v := range myAnimal{
	// 	v.Eat()
	// }

}