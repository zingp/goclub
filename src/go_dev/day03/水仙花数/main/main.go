package main
/*
找出100-1000之间的水仙花数
水仙花数：个位的三次方+十位的三次方+百位的三次方 = 本身
*/
import "fmt"

func main() {
	for i:=100;i<1001;i++ {
		b := 0
		for j:=i;j>0;j/=10 {
			b += (j%10) * (j%10) * (j%10)
		}
		if b == i{
			fmt.Printf("%d 是水仙花数\n", i)
		}
	}
}

