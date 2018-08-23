package multiArray

import "fmt"

func TwoArray()  {
	// 声明一个二维整型数组，两个维度分别存储 4 个元素和 2 个元素
	var array [4][2]int

	// 使用数组字面量来声明并初始化一个二维整型数组
	array1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化外层数组中索引为 1 个和 3 的元素
	array2 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化外层数组和内层数组的单个元素
	array3 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

}
