package main
/*计算阶乘之和*/

import "fmt"

func sum(n int) uint64 {
	/*计算阶乘之和*/
	var s uint64 = 1
	var sum uint64 = 0
	for i:=1;i<=n;i++ {
		s = s * uint64(i)
		sum += s
	}
	return sum
}

func main() {
	var n int
	fmt.Println("请输入如一个正整数：")
	fmt.Scanln(&n)
	s:= sum(n)
	fmt.Printf("n!+...1！= %d", s)
}
