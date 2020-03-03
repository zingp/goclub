package main

import "fmt"
// 一个struct通过嵌套匿名结构体可实现继承
//struct嵌套有名结构体叫组合
type Car struct {
	Name string
	weight int
}

func (this *Car) run() {
	fmt.Printf("The %s is runing.\n",this.Name)
}

type train struct {
	Car   // 这里实现了继承
	Node int
}

func main()  {
	var t train
	t.Name = "myTrain"
	t.weight = 100000
	t.Node = 7

	fmt.Println(t)
	t.run()
}