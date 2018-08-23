package main

import (
	"fmt"
)

func justify (items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%d个参数是bool\n", i)
		case int32:
			fmt.Printf("第%d个参数是int32\n", i)
		case string:
			fmt.Printf("第%d个参数是string\n", i)
		}
	}
}

func main() {
	var a int32
	var b string
	var c bool
	justify(a, b, c)
}