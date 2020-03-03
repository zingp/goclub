package statMap

import "fmt"

func MakeStatMap() {

	dict := make(map[string]string)
	fmt.Println(dict)
}

// 创建一个映射，键和值的类型都是 string
// 使用两个键值对初始化映射
func WordStatMap() {
	dict := map[string]string{"name": "liuyouyuan", "age": "23", "sex": "male"}
	fmt.Println(dict)
}

// 使用映射字面量声明空映射
func EmptyMap()  {
	dict := map[string]int{}
	fmt.Println(dict)
}