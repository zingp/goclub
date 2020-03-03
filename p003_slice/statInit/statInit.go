package statInit

import "fmt"

// 通过make声明切片
func MakeStatSlice() {
	//使用长度声明一个字符串切片
	// 创建一个字符串切片
	// 其长度和容量都是 5 个元素
	slice := make([]string, 5)
	slice[1] = "liu"
	slice[2]  = "you"

	fmt.Println(slice)

	//使用长度和容量声明整型切片
	slice2 := make([]int, 3, 5)
	for i:=0; i<3; i++ {
		slice2[i] = i
	}

	//不允许创建容量小于长度的切片
	fmt.Println(slice2)
}


//通过切片字面量来声明切片
func WordStatSlice()  {
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	slice1 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice1)

	// 创建一个整型切片
	// 其长度和容量都是 3 个元素
	slice2 := []int{10, 20, 30}
	fmt.Println(slice2)


	// 创建字符串切片
	// 使用空字符串初始化第 100 个元素
	slice3 := []string{99: ""}
	fmt.Println(slice3)
	//如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。
	//只有不指定值的时候，才会创建切片
}


// 切片和数组不同
func DiffArraySlice()  {
	// 创建有 3 个元素的整型数组
	array := [3]int{10, 20, 30}
	// 创建长度和容量都是 3 的整型切片
	slice := []int{10, 20, 30}

	fmt.Println("数组", array)
	fmt.Println("切片", slice)
}

// nil 和空切片
func EmptySlice() {

	// 创建 nil切片
	// 只要在声明时不做任何初始化，就会创建一个 nil 切片
	// 创建 nil 整型切片 长度0 容量0
	var slice []int
	fmt.Println("nil 切片:", slice)

	// 声明空切片
	emptySlice := make([]int, 0)
	fmt.Println("空切片", emptySlice)

	emptySlice2 := []int{}
	fmt.Println("空切片2", emptySlice2)

}

// 多维slice
func MultiSlice() {
	slice := [][]int{{10, 20}, {30, 40, 50}}
	slice[0] = append(slice[0], 30)
	fmt.Println("老多维切片", slice)
	fmt.Println("新多维切片", slice[0])
}
