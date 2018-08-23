package main
/*
打印：
A
AA
AAA
AAAA
AAAAA
*/
import "fmt"

func main() {
	for i:=1;i<6;i++{
		for j:=0;j<i;j++ {
			fmt.Printf("%s","A")
		}
		fmt.Println()
	}
}
