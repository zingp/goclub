/*输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数*/
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("输入一个字符串：")
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)      //如果用Scanln获取不到空格
	s,err := inputReader.ReadString('\n')    //获取键盘输入，包括空格
	if err != nil {
		fmt.Printf("error: %s",err)
		return
	}

	si := []byte(s)
	var sumAbc, sumNum, sumSpace, sumOther int
	fmt.Println(si)
	for i:=0;i<len(si);i++ {
		if si[i] == 32 {
			sumSpace += 1
		}else if 48 <= si[i] && si[i] <= 57 {
            sumNum += 1
		}else if 65 <= si[i] && si[i] <= 90 {
			sumAbc += 1
		}else if 97 <= si[i] && si[i] <= 122 {
			sumAbc += 1
		}else {
			sumOther += 1
		}
	}

	fmt.Printf("数字%d个，字母%d个，空格%d个，其他字符%d个。",sumNum,sumAbc,sumSpace,sumOther)

}
