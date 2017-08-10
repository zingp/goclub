package main

import "fmt"

/*
map的声明方式：

*/

func testMap() {
	var m map[string]string
	m = make(map[string]string)
	m["name"] = "lyy"
	fmt.Println(m)
}

func testMap2(){
	a := make(map[string]map[string]string)
	a["name"] = make(map[string]string)
	a["name"]["first"] = "Liu"
	a["name"]["second"] = "You"
	a["name"]["third"] = "Yuan"
	fmt.Println(a)

	//字典的遍历
	for k,v := range a {
		fmt.Println(k, "---", v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}

func main() {
	testMap()
	testMap2()
}