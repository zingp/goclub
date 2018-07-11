package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
)

//方法一
func ReadFile(d string){

	file, err := os.Open(d)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close()

	var data [1024]byte
	for {
		n, err := file.Read(data[:])

		// 如果读完则终止循环
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file error:", err)
		}

		str := string(data[0:n])
		fmt.Println(str)
	}
}

//方法二：通过bufio读取文件
func ReadFlieBufio(d string) {
	file, err := os.Open(d)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
		}
		fmt.Println(line)
	}

}

//方法三:一次性读取整个文件，别用来读取大文件
func testIoutilReadFile(d string) {
	dataBytes, err := ioutil.ReadFile(d)
	if err != nil {
		fmt.Println("read file err:", err)
	}
	fmt.Println(string(dataBytes))

}

func main() {
	var dir string = "C:/Users/liuyouyuan/Desktop/笔记/部署文档/centos6下go语言环境搭建.txt"
	//ReadFile(dir)
	//ReadFlieBufio(dir)
	testIoutilReadFile(dir)
}
