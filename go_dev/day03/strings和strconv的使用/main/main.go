package main
/*
HasPrefix(s, prefix string) bool 判断字符串s是否以prefix开头；
HasSuffix(s, suffix string) bool 判断字符串s是否以prefix结尾。
strings.Index(s, sep string) int 判断sep在s中首次出现的位置，没有则返回-1；
strings.LastIndex(s, sep string) int判断sep在s中最后出现的位置，没有则返回-1；
*/
import (
	"strings"
	"fmt"
)

func main() {
	str := "http://www.cnblogs.com/zingp/"
	r1 := strings.HasPrefix(str, "http://")
	fmt.Println(r1)
	r2 := strings.HasSuffix(str, "/")
	fmt.Println(r2)

	r3 := strings.Index(str, "c")
	fmt.Println(r3)
	r4 := strings.LastIndex(str, "c")
	fmt.Println(r4)
}
