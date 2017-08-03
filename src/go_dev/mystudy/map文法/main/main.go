package main
/*
map的文法,
如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
*/

import "fmt"

type PersonInfo struct{
	Name,Sex string
	Age int
}

var m map[string]string

func main() {
	m = map[string]string{
		"name": "liuyouyuan",
		"age":"25",
		"sex":"male",
	}
	fmt.Println(m)

	person := map[int]PersonInfo{
		1001: {
			Name:"liuYouYuan",
			Age:25,
			Sex:"male",
		},
		1002:{
			Name:"zing-p",
			Age:26,
			Sex:"male",
		},
	}
	fmt.Println(person)
}
/*
result:
map[name:liuyouyuan age:25 sex:male]
map[1001:{liuYouYuan male 25} 1002:{zing-p male 26}]
*/