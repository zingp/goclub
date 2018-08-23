/*
输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
左读完全相同的字符串。
此处的做法就是先把字符串反转，对比
*/
package main

import "fmt"

func isHW(s string) bool {
	var result []byte
	li := []byte(s)
	for i:=0;i<len(s);i++ {
		result = append(result, li[len(s)-i -1])
	}
	// 将反转后的slice转换成string再和原来的s对比
	if string(result) == s {
		return true
	}
	return false
}

func main() {
	var s string
	fmt.Println("输入一个字符串：")
	fmt.Scanln(&s)
	if isHW(s) {
		fmt.Printf("%s is huiwen...",s)
	}else {
		fmt.Printf("%s is not huiwen...",s)
	}
}
