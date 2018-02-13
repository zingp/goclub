package useSlice

import "fmt"

// 赋值和切片
// 对切片里某个索引指向的元素赋值和对数组里某个索引指向的元素赋值的方法完全一样。
// 使 用[]操作符就可以改变某个元素的值
func AssignmentAndSlice(){
	// 使用切片字面量来声明切片
	// 创建一个整型切片
	// 其容量和长度都是 5 个元素
	slice := []int{1, 2, 3, 4, 5}
	slice[0] = 6
	fmt.Println(slice)

	// 使用切片创建切片
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1: 3]
	fmt.Println("切片创建切片", newSlice)
	/*第一个切片 slice 能够看到底层数组全部 5 个元素的容量，不过之后的 newSlice 就看不
到。对于 newSlice， 底层数组的容量只有 4 个元素。 newSlice 无法访问到它所指向的底层数
组的第一个元素之前的部分。所以，对 newSlice 来说，之前的那些元素就是不存在的。*/
	/*
	对底层数组容量是 k 的切片 slice[i:j]来说
		长度: j - i
		容量: k - i
	*/

	// 修改切片内容可能导致的结果
	// 现在两个切片共享同一个底层数组。如果一个切片修改了该底层数组的共享
	// 部分，另一个切片也能感知到
	newSlice[1] = 88
	fmt.Println("修改newSlice后的slice", slice)
	fmt.Println("修改newSlice后的newSlice", newSlice)
}

// 切片增长
func AppendSlice() {
	slice := []int{10 ,20, 30, 40, 50}

	// 使用 append 增加切片的长度
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println("slice>>:", slice)
	fmt.Println("newSlice>>:", newSlice)

	// 使用 append 同时增加切片的长度和容量
	newSlice2 := append(slice, 60)
	fmt.Println(slice)
	fmt.Println(newSlice2)

	/*当这个 append 操作完成后， newSlice 拥有一个全新的底层数组，这个数组的容量是原来的两倍
	函数 append 会智能地处理底层数组的容量增长。在切片的容量小于 1000 个元素时，总是
	会成倍地增加容量。一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25%
	的容量。随着语言的演化，这种增长算法可能会有所改变。
	*/

}

// 创建切片时的 3 个索引
func ThreeIndexBuildSlice() {
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 将第三个元素切片，并限制容量
	// 其长度为 1 个元素，容量为 2 个元素
	fmt.Println(source)
	slice := source[2:3:4]
	fmt.Println("使用第三个限制容量：", slice)
	// 如何计算长度和容量
	// 对于 slice[i:j:k] 或 [2:3:4]长度: j – i 或 3 - 2 = 1
	// 容量: k – i 或 4 - 2 = 2
}

func ThreeIndexSetLenEqBound() {
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是 1 个元素
	slice := source[2:3:3]
	// 向 slice 追加新字符串
	slice = append(slice, "Kiwi")
	fmt.Println(source)
	fmt.Println(slice)
}


// 将一个切片追加到另一个切片
func SliAppendSli() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	//slice2 = append(slice1, slice2...)

	//fmt.Println("追加切片之后的slice2", slice2)
	fmt.Printf("slice2追加到slice1之后：%v\n", append(slice1, slice2...))
}

// 迭代切片
func IterSlice()  {
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}

	for index, v := range slice {
		fmt.Println(index, v)
	}

}
