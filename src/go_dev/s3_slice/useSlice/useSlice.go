package useSlice

import "fmt"

// 赋值和切片
// 对切片里某个索引指向的元素赋值和对数组里某个索引指向的元素赋值的方法完全一样。
// 使 用[]操作符就可以改变某个元素的值
func AssignmentAndSlice(){
	// 创建一个整型切片
	// 其容量和长度都是 5 个元素
	slice := []int{1, 2, 3, 4, 5}
	slice[0] = 6
	fmt.Println(slice)
}
