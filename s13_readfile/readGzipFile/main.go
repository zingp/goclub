package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)
// gzip 打开文件
func testReadGzip(d string) {
	// [1] 用os.Open打开文件，获得一个句柄
	f, err := os.Open(d)

	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	 
	// [2] gzip.NewReader(文件句柄)
	gzReader, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println("read gzip file err:", err)
		return
	}

    // [3] bufio.NewReader()
	bufReader := bufio.NewReader(gzReader)
	for {
		line, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ReadString err:", err)
			return
		}
		fmt.Println(line)
	}
}

func main() {
	fileName := "C:/Users/liuyouyuan/Desktop/笔记/部署文档/centos7 安装mysql57.txt.gz"
	testReadGzip(fileName)
}
