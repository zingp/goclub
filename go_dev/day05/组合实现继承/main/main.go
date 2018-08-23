package main

import "fmt"
// 一个struct通过嵌套匿名结构体可实现继承
//struct嵌套有名结构体叫组合
//继承通过组合实现
type Car struct {
	Name string
	weight int
}

func (this *Car) run() {
	fmt.Printf("The %s is runing.\n",this.Name)
}

type train struct {
	c Car   // 这里实现了继承
	Node int
}

func main()  {
	var t train
	t.c.Name = "myTrain"   //注意访问方式的改变
	t.c.weight = 100000
	t.Node = 7

	fmt.Println(t)
	t.c.run()
}
