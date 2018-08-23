/*
一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
编程找出1000以内的所有完数。
*/

package main

import "fmt"

func isWS(n int) bool{
	sum := 1
	for i:=2;i*i<=n;i++ {
		if n%i == 0{
			sum += i
			sum += n/i
		}
	}
	if sum == n {
		return true
	}
	return false
}

func main() {
	for i:=1; i<=1000;i++ {
		if isWS(i) {
			fmt.Printf("%d is wanshu...\n", i)
		}
	}
}
