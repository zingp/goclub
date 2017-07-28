package main

import (
	"fmt"
)
/*
单引号表示单个字符对应的ASCII码的十进制；
反引号里面是什么内容就是什么内容；
双引号会打印换行等，可以格式化输出。
*/
func main()  {

	var cha  byte = 'b'
	var c = 'a'
	var str  string = `落霞与孤鹜齐飞，\n秋水共长天一色`
	var strDouble  string = "落霞与孤鹜齐飞，\n秋水共长天一色"

	fmt.Println(cha)
	fmt.Println(c)
	fmt.Println(str)
	fmt.Println(strDouble)
}

/*
代码运行结果：
98
落霞与孤鹜齐飞，\n秋水共长天一色
落霞与孤鹜齐飞，
秋水共长天一色
*/
