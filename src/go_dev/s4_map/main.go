package main

import (
	"go_dev/s4_map/statMap"
	"go_dev/s4_map/useMap"
)

func main() {
	statMap.MakeStatMap()
	statMap.WordStatMap()
	statMap.EmptyMap()

	// 从映射获取值并判断键是否存在
	useMap.MapValueExist()

	// map的遍历
	useMap.RangeMap()

	// 删除key value
	useMap.DeleteKeyValue()

	// 函数间传递map
	useMap.TransferMapInFunc()
}
/*
小结：
	数组是构造切片和映射的基石。
	Go 语言里切片经常用来处理数据的集合，映射用来处理具有键值对结构的数据。
	内置函数 make 可以创建切片和映射，并指定原始的长度和容量。也可以直接使用切片
和映射字面量，或者使用字面量作为变量的初始值。
	切片有容量限制，不过可以使用内置的 append 函数扩展容量。
	映射的增长没有容量或者任何限制。
	内置函数 len 可以用来获取切片或者映射的长度。
	内置函数 cap 只能用于切片。
	通过组合，可以创建多维数组和多维切片。也可以使用切片或者其他映射作为映射的值。
但是切片不能用作映射的键。
	将切片或者映射传递给函数成本很小，并且不会复制底层的数据结构
 */
