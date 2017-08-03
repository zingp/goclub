package main
//一个函数和与其相关的引用环境组合而成的实体
import "fmt"

func main()  {
	var f = Adder()   //返回值是个函数
	fmt.Println(f(1), " _ ")
	fmt.Println(f(20), " _ ")
	fmt.Println(f(300))

}

//定义一个Adder()函数，没有参数；它的返回值是一个函数func(int)
//func(int)的返回值是int
func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
