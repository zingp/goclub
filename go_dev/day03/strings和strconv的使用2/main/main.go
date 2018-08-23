package main
/*
strings.Replace(s, old, new string, n int) string 用new替换s中的old字符串n次；
strings.Count(s, sep string) int 统计sep在s中出现的次数。
strings.Repeat(s string, count int) string 重复count次s。
*/

import (
	"strings"
	"fmt"
)

func main() {
	str := "http://www.cnblogs.com/zingp/"

	r1 := strings.Replace(str, "www", "666", 1)
	fmt.Println(r1)
	r2 := strings.Count(str, "p")
	fmt.Println(r2)

	str2 := "str"
	r3 := strings.Repeat(str2, 3)
	fmt.Println(r3)

	r4 := strings.ToLower("LIU")  //转换为小写
	fmt.Println(r4)
	r5 := strings.ToUpper("youyuan")  //转换为大写
	fmt.Println(r5)
}
