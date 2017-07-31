package main
/*判断素数和获取用户输入*/
import (
	"fmt"
)

func isPrime(n int) bool{
	if n==1 || n==2{
		return true
	}
	for i:=2; i*i<n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("输入：")
	fmt.Scanln(&n)
	res := isPrime(n)
	if res {
		fmt.Printf("%d 是素数", n)
	}else {
		fmt.Printf("%d 不是素数", n)
	}
}
