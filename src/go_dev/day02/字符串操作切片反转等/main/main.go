package main

import "fmt"

func reverse(str string) string {
	/*反转字符串一*/
	var resStr string
	strLen := len(str)
	for i:=0; i<strLen; i++{
		resStr =  resStr + fmt.Sprintf("%c", str[strLen - i -1])
	}
	return resStr
}

func reverse1(str string) string {
	/*反转字符串二*/
	var result []byte
	tmp := []byte(str)
	strLen := len(str)
	for i:=0;i<strLen;i++ {
		result = append(result, tmp[strLen-i-1])
	}
	return string(result)

}

func main() {
	a := "Hello"
	b := "World"

	str1 := a + " " + b
	fmt.Println(str1)
	fmt.Printf("%s %s\n", a, b)

	fmt.Println(a[2:5])  //[2,5)
	fmt.Println("长度", len(b))  //len的算法是o(1)的，默认保存起来了

	str2 := reverse(str1)
	fmt.Printf("第一次反转结果：%s\n",str2)
	str3 := reverse1(str2)
	fmt.Printf("第二次反转结果：%s\n",str3)
}

/*
运行结果：
Hello World
Hello World
llo
长度 5
第一次反转结果：dlroW olleH
第二次反转结果：Hello World
*/
