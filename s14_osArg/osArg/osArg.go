package main

import (
	"flag"
	"os"
	"fmt"
)

func testOsArgs(){
	// os.Args is a slice string
	args := os.Args
	// 0 是二进制本身 1是二进制后的第1个参数 ...
	for i,v := range args {
		fmt.Printf("args[%d]=%s\n",i, v)
	}
}


var (
	file string
	row int
)
func init(){
	// go run osArg.go -c etc/myconf -n 20
	// 定义的是c命令行用-c,定义的是-c，命令行用--c
	flag.StringVar(&file, "c", "d:/test.cfg", "请指定文件")
	flag.IntVar(&row, "n", 10, "请指定行数")
	flag.Parse()  // 必须要解析
}

func main() {
	// testOsArgs()
	fmt.Printf("file=%s; row=%d", file, row)
	
}