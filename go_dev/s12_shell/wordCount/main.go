package main
// 统计终端输入的字符串中的英文字母、空格、数字、其他字符个数

import (
	"fmt"
	"bufio"
	"os"
)

func WordCount(s string){

	var enCount, spCount, numCount, otherCount int
	utf8Str := []rune(s)

	for _, v := range utf8Str {
		if v >= 'a' && v <= 'z' ||  v >= 'A' && v <= 'Z'{
			enCount++
			continue
		}
		if v == ' ' {
			spCount++
			continue
		}
		if v >='0' && v <= '9'{
			numCount++
			continue
		}
		otherCount++
	}

	fmt.Printf("enCount=%d, spCount=%d, numCount=%d, otherCount=%d\n",
		enCount, spCount, numCount, otherCount)
}

func ReadInput() {
	reader := bufio.NewReader(os.Stdin)
	//line, _:= reader.ReadString('\n')
	line, _, _:= reader.ReadLine()
	WordCount(string(line))
}

func main() {
	ReadInput()
}
