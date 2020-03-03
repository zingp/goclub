package main
/*
多维数组，及其遍历；
注意v1也是一个数组
*/

import (
	"fmt"
)

func main() {
	var a [2][5]int = [...][5]int{{1,2,3,4,5},{6,7,8,9,10}}

	for row,v1 := range a {
		for col,v2 := range v1 {
			fmt.Printf("[%d][%d]== %d\n",row,col,v2)
		}

	}
}
