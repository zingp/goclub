package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "我们 0.242507 -0.144510 0.401783 -0.471121 "
	// 返回子串sep在字符串s中第一次出现的索引值，不在的话返回-1.
	idx := strings.Index(line, " ")
	// chars中任何一个Unicode代码点在s中首次出现的位置，不存在返回-1
	idx2 := strings.IndexAny(line, " ")

	fmt.Println(line[:idx2], idx2)
	fmt.Println(idx)

	s := "AG 输了"
	fmt.Println(strings.ToLower(s))
}
