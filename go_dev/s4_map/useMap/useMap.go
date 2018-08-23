package useMap

import "fmt"

//从映射获取值并判断键是否存在
func MapValueExist()  {
	personInfo := map[string]string{"name": "liuyouyuan", "age": "23", "sex": "male"}

	value, exist := personInfo["name"]
	if exist {
		fmt.Println("person name:", value)
	}

	// 方法2
	// 这个键存在吗？
	if value != "" {
		fmt.Println(value)
	}
	/*在 Go 语言里，通过键来索引映射时，即便这个键不存在也总会返回一个值。在这种情况下，
返回的是该值对应的类型的零值。*/
}

func RangeMap()  {
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral": "#ff7F50",
		"DarkGray": "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func DeleteKeyValue() {

	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral": "#ff7F50",
		"DarkGray": "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// 删除键为 Coral 的键值对
	// 这种方法只能用在映射存储的值都是非零值的情况
	delete(colors, "Coral")
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s <--> Value: %s\n", key, value)
	}
}

func TransferMapInFunc() {
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral": "#ff7F50",
		"DarkGray": "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Old Key: %s Value: %s\n", key, value)
	}
	// 调用函数来移除指定的键
	removeColor(colors, "Coral")
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("After delete Key: %s Value: %s\n", key, value)
	}
}

// removeColor 将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
/*
在函数间传递映射并不会制造出该映射的一个副本。实际上，当传递映射给一个函数，并对
这个映射做了修改时，所有对这个映射的引用都会察觉到这个修改
*/
