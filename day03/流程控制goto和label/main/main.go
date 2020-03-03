package main
/* goto 感觉可读性比较差 */

import "fmt"

func test(){

	for i:=0; i<=5;i++ {
		HHH:
		if 2 <i && i < 5{
			fmt.Printf("test i is %d\n", i)
			break
		}
		if i == 3{
			goto HHH
		}
		fmt.Printf("test:::i = %d\n", i)

	}
}

func main() {
	LABEL1:
		for i:= 0; i <= 5; i++ {
			for j:=0; j<=5; j++ {
				if j ==4 {
					continue LABEL1
				}
				fmt.Printf(" i is :%d, and j is :%d\n", i, j)
			}
		}
	test()
}
