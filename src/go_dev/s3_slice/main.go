package main

import (
	"go_dev/s3_slice/statInit"
	"go_dev/s3_slice/useSlice"
)

func main() {
	// 用make申明切片
	statInit.MakeStatSlice()

	// 用字面量申明切片
	statInit.WordStatSlice()

	// 数组和切片的不同
	statInit.DiffArraySlice()

	// nil 和空切片
	statInit.EmptySlice()

	// 切片的使用
	// 赋值和切片
	useSlice.AssignmentAndSlice()

	// 切片的增长
	useSlice.AppendSlice()
}