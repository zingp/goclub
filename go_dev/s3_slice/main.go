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

	// 多维切片
	statInit.MultiSlice()

	// 切片的使用
	// 赋值和切片
	useSlice.AssignmentAndSlice()

	// 切片的增长
	useSlice.AppendSlice()

	// 三个索引创建切片
	useSlice.ThreeIndexBuildSlice()

	// 三个索引控制新切片长度和容量一致
	useSlice.ThreeIndexSetLenEqBound()

	// 切片追加到切片
	useSlice.SliAppendSli()

	// 迭代切片
	useSlice.IterSlice()

	// range 提供了每个元素的副本
	useSlice.FactRange()

	// for 遍历slice
	useSlice.ForSlice()
}

