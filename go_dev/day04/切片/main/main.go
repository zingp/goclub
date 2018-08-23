package main

import "fmt"

/*
切片是数组的一个引用，因此，切片是引用类型；
切片长度可变，因此，切片是一个可变数组；
切片的遍历方式和数组一样，用len()求长度；
cap()可以求出slice最大的容量：0<=len(slice)<=cap(array),其中array是slice引用的数组；
切片的定义：var变量名 []类型，比如：var str []string; var arr []int
切片的内存布局是：
ptr ->数组；len;cap;三个部分；
通过make来创建切片：var slice []type = make([]type, len)
				 slice := make([]type, len)
				 slice := make([]type, len, cap)
				 自己会维护一个数组
*/


func modify(a []int) {
	a[1] = 100
}

func main() {
	var b []int = []int{1,2,3,4,5}   //通过数组来创建
	modify(b)
	fmt.Printf("b = %d\n", b)

	var mySlice []int = make([]int, 10, 20)
	mySlice[0] = 10
	mySlice[9] = 100
	fmt.Printf("mySlice= %d\n len= %d\n cap= %d\n",mySlice,len(mySlice),cap(mySlice))

	sli := make([]string, 5)
	sli[0] = "lyy"
	sli[1] = "zing-p"
	fmt.Println(sli,len(sli),cap(sli))

	fmt.Println("addr:", &sli[0])    //扩容之前的地址

	sli = append(sli,"8")
	sli = append(sli,"8")
	sli = append(sli,"8")

	fmt.Println(sli)
	fmt.Println(len(sli),cap(sli))
	fmt.Println("addr:", &sli[0])   //扩容之后地址改变

	sli = append(sli,"9")
	sli = append(sli,"9")
	sli = append(sli,"9")

	fmt.Println(sli)
	fmt.Println(len(sli),cap(sli))      //扩容：5 -> 10 -> 20
}