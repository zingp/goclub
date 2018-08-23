package main

import (
	"math/rand"
	"fmt"
)

func main() {
	var (
		n int
		a int
	)
	n = rand.Intn(100)

	for {
		fmt.Println("输入你猜的数字：")
		fmt.Scanf("%d\n", &a)
		flag := false
		switch  {
		case a == n:
			fmt.Println("猜对了。")
			flag = true
		case a > n:
			fmt.Println("猜大了。")
		case a < n:
			fmt.Println("猜小了。")
		}
		if flag{
			break
		}
	}
}
