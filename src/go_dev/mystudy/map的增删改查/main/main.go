package main

import "fmt"

func main() {
	m := map[string]string{
		"001":"lyy",
	}
	m["002"] = "liuyouyuan"
	m["003"] = "zing-p"
	fmt.Println(m)

	delete(m, "002")
	fmt.Println(m)

	ele := m["003"]
	fmt.Println(ele)

	/*如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
    同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。*/
	elem, ok := m["002"]
	fmt.Println(elem, ok)
}

/*
result:
map[001:lyy 002:liuyouyuan 003:zing-p]
map[001:lyy 003:zing-p]
zing-p
 false   //注意这前面有个空字符串。
*/
