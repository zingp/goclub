package main

import(
	"os"
	"fmt"
	"bufio"
)

// 写文件 方法一
func testWriteFile(d string){
	f, err := os.OpenFile(d, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	defer f.Close()

	for i:=0;i<10;i++ {
		f.WriteString(fmt.Sprintf("i am studying go program %d\n",i))
	}
}


//写文件方法二：用bufio实现
func writeFileBufio(d string) {
	f, err := os.OpenFile(d, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	defer f.Close()

	bufWriter := bufio.NewWriter(f)
	for i:=0;i<10;i++ {
		bufWriter.WriteString(fmt.Sprintf("Write file by bufio %d\n",i))
	}
	bufWriter.Flush()
	f.Seek(0, 2)
	f.WriteString("80000 200\n")
}

func main() {
	
	writeFileBufio("./test.txt")
}
