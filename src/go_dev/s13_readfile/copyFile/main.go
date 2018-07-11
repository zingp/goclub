package main

import(
	"os"
	"io"
	"fmt"
)


// 文件拷贝：目标文件  源文件【新文件在前，旧文件在后】
func CopyFile(dstFile string, srcFile string)(writen int64, err error) {
	f1, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	defer f1.Close()

	f2, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	defer f2.Close()

	return io.Copy(f2, f1)
}

func main(){
	CopyFile("D:/goCopyLog.txt", "D:/goTestLog.txt")
}