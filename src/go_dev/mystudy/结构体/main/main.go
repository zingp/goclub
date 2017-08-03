package main
/*
一个结构体（`struct`）就是一个字段的集合
*/
import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	v := Point{2,3}  //赋值
	fmt.Println(v)
	v.x = 6           //结构体中的成员赋值
	fmt.Println(v)

	p := &v
	p.x = 1e9
	p.y = 100
	fmt.Println("结构体：", v)
}

