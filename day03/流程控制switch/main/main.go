package main

import "fmt"
/*和其他语言的switch的语法的区别是用break*/

func main() {
	var a int = 100
	switch a {
	case 0:
		fmt.Println("a is equal 0.")
		//fallthrough  穿透
	case 10, 100:  //多个条件可以写一行
		fmt.Println("a is equal 10 or 100")
	default:
		fmt.Println("a is equal default.")
	}

	switch  {
	case a > 0 && a < 50:
		fmt.Println("a < 50.")
	case a > 50 && a <= 100:
		fmt.Println("50 < a <= 100.")
	default:
		fmt.Println("i don not kwon.")
	}
}
