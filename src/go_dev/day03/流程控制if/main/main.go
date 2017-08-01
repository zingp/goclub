package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Println("please input number or alphabet:")
	fmt.Scanln(&a)

	number, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("%s can not convert to int.err:%s", a, err)
		return
	}else {
		fmt.Println(number)
	}
}
