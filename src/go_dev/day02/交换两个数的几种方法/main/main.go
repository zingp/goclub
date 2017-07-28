package main

import "fmt"

/* 交换两个变量的值的三种方法*/


func change(a *int ,b *int)  {
	temp := *a
	*a = *b
	*b = temp
}

func change2(a int, b int)(int, int) {
	return b, a
}

func main()  {
	first := 100
	second := 200
	/*方法一*/
	//change(&first, &second)

	/*方法二*/
	//first, second = change2(first,second)

	/*方法三*/
	first, second = second, first

	fmt.Println("first=", first)
	fmt.Println("second=", second)

}