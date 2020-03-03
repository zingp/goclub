package testadd

import "fmt"

func init()  {
	fmt.Println("in testadd package init func...")
}

var Name = "Liu You."
var Age int = 25

func AddFunc(a int, b int) int{
	return a+b
}
