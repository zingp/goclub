package main
/*字符串转int,根据ASCII码转换来计算，很巧*/

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	result := 0
	for i:=0;i<len(str);i++ {
		num := int(str[i] - '0')
		result += (num * num * num)
	}
	number, err := strconv.Atoi(str)
	if (err != nil) {
		fmt.Printf("can not convert %s to int.", str)
		return
	}

	if result == number {
		fmt.Printf("%d is 水仙花数", number)
	}else {
		fmt.Printf("%d is not 水仙花数", number)
	}

}
