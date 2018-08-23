package useArray

import (
	arr "go_dev/s2_array/statmentAndInit"
	"fmt"
)

// 修改索引为 2 的元素的值
func ModArray()  {
	arr.Array[2] = 35
	fmt.Println(arr.Array)
}


func PointArr()  {
	// 声明包含 5 个元素的指向整数的数组
	// 用整型指针初始化索引为 0 和 1 的数组元素
	array := [5]*int{0: new(int), 1: new(int)}
	// 为索引为 0 和 1 的元素赋值
	*array[0] = 10
	*array[1] = 20
	fmt.Println("指针数组", array)
	for i:=0;i < len(array); i++ {
		fmt.Println("指针数组的值", *array[1])
	}

	//指针默认为nil
	//使用int指针初始化这个数组
	array2 := [5]*int{0: new(int), 1: new(int), 2: new(int), 3: new(int), 4: new(int)}
	*array2[4] = 40
	for i:=0;i < len(array2); i++ {
		fmt.Println("array2指针数组的值", *array2[i])
	}
}

//编译器会阻止类型不同的数组互相赋值
/*
func ErrPointArrCopy()  {

	// 声明第一个包含 4 个元素的字符串数组
	var array1 [4]string
	// 声明第二个包含 5 个元素的字符串数组
	// 使用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 将 array2 复制给 array1
	array1 = array2

	fmt.Println(array1)

}
*/

//把一个指针数组赋值给另一个
func PointArrCopy()  {
	// 声明第一个包含 3 个元素的指向字符串的指针数组
	var array1 [3]*string
	// 声明第二个包含 3 个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array2 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	// 将 array2 复制给 array1
	array1 = array2
	for i := 0; i < len(array1); i++ {
		fmt.Println(*array1[i])
	}
}