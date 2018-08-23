package main

import(
	"fmt"
	"go_dev/day02/初始化函数/testadd"
)
// init函数会在main函数之前执行，常用作初始化。
func init()  {
	fmt.Println("do something about init...")
}

func main()  {
	a := 10
	b := 2
	c := testadd.AddFunc(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, c)

	fmt.Println("name:",testadd.Name)
	fmt.Println("age:", testadd.Age)
}
