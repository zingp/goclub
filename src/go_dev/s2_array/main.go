package main



import (
	arr "go_dev/s2_array/statmentAndInit"
	"fmt"
	arr_use "go_dev/s2_array/useArray"
	arr_mul "go_dev/s2_array/multiArray"
)

func main() {
	fmt.Println(arr.Array)
	fmt.Println(arr.Array2)
	fmt.Println(arr.Array3)
	fmt.Println(arr.Array4)

	arr_use.ModArray()

	arr_use.PointArr()

	//编译器会阻止类型不同的数组互相赋值
	//arr_use.ErrPointArrCopy()

	//把一个指针数组赋值给另一个
	arr_use.PointArrCopy()

	//二维数组
	arr_mul.TwoArray()

	minLen := 1 << 30
	fmt.Println(minLen)

}
