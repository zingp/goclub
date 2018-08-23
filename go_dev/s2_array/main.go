package main

import "fmt"

func main() {
	var array = [...]int{1, 2, 3, 4, 5}

	// 修改元素
	array[2] = 10

	// 查询元素
	fmt.Println("访问数组索引为2的元素", array[2])    // 访问数组索引为2的元素 10

	//访问指针数组
	array2 := [5]*int{0: new(int), 1: new(int)}
	fmt.Println("指针数组元素2", array2[2])          // 指针数组元素2 <nil>

}

//import (
//	arr "go_dev/s2_array/statmentAndInit"
//	"fmt"
//	arr_use "go_dev/s2_array/useArray"
//	arr_mul "go_dev/s2_array/multiArray"
//)

//func main() {
//	fmt.Println(arr.Array)
//	fmt.Println(arr.Array2)
//	fmt.Println(arr.Array3)
//	fmt.Println(arr.Array4)
//
//	arr_use.ModArray()
//
//	arr_use.PointArr()
//
//	//编译器会阻止类型不同的数组互相赋值
//	//arr_use.ErrPointArrCopy()
//
//	//把一个指针数组赋值给另一个
//	arr_use.PointArrCopy()
//
//	//二维数组
//	arr_mul.TwoArray()
//
//	minLen := 1 << 30
//	fmt.Println(minLen)
//
//}
