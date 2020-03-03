package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	str := " hello world "
	fmt.Println(str)
	r1 := strings.TrimSpace(str) //去掉字符串两端的空格
	fmt.Println(r1)

	str2 := "i am i"
	r2 := strings.Trim(str2,"i")
	fmt.Println(r2)

	str3 := "helllo"
	r3 := strings.TrimLeft(str3, "he")   //删除开头
	r4 := strings.TrimRight(str3, "lo")  //删除末尾，但是同样的字符会被删除
	fmt.Println(r3)
	fmt.Println(r4)

	st := "I love beijing tiananmen."
	r5 := strings.Fields(st)   //返回空格分割的所有字串的数组
	fmt.Println(r5)
	fmt.Println(r5[0],r5[2])

	s := "name:lyy|age:25"
	r6 := strings.Split(s, "|")   //返回以指定分割符分割的所有子串
	fmt.Println(r6)
	r7 := strings.Split(r6[0], ":")
	fmt.Println(r7)

	r8 := strings.Join(r7, "+")
	fmt.Printf("r8 is: %s\n", r8)

	var a int = 888
	str_a := strconv.Itoa(a)   //将整数转换为字符串，证书必须为int类型，int32什么的不行
	fmt.Printf("str_a is %s\n", str_a)

	num := "999"
	number, err := strconv.Atoi(num)    //将字符串转换为数字
	if err != nil {
		fmt.Printf("error:%s\n",err)
		return
	}
	fmt.Println(number)
}
