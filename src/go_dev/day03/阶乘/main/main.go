package main
/*计算阶乘之和*/

import "fmt"

func jiechen(n int) int {
	/*计算阶乘*/
	sum := 1
	for i:=1;i<=n;i++ {
		sum *= i
	}
	return sum

}

func main() {
	var n int
	fmt.Println("请输入如一个正整数：")
	fmt.Scanln(&n)
	sum := 0
	for i:=1;i<=n;i++ {
		sum += jiechen(i)
	}
	fmt.Printf("n!+...1！= %d", sum)
}
