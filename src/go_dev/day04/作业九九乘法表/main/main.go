/*
打印99乘法表
*/
package main

import "fmt"

func printMultiply(){
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%2d ",i,j, i*j)
		}
		fmt.Print("\n")
	}
}

func main() {
	printMultiply()
}
