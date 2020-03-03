package main

import (
	"fmt"
	"strings"
)


func makeSuffixFunc(suffix string)  func(st string) string{
	return func(name string) string{
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}
func main()  {
	f1 := makeSuffixFunc(".bmp")
	f2 := makeSuffixFunc(".jpg")
	fmt.Println(f1("test"))
	fmt.Println(f2("test"))

}
