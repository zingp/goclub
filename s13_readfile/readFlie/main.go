package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
	"time"
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
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Println("read line err", err)
			continue
		}
		//fmt.Println(line)
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

// 按照行读，ReadLine会更快
func bufioReadLine(d string) {
	file, err := os.Open(d)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0
	for {
		//line, err := reader.ReadString('\n')
		line, isprefix, err := reader.ReadLine()
		for isprefix && err == nil {
			var bs []byte
			bs, isprefix, err = reader.ReadLine()
			line = append(line, bs...)
		}
		if err == io.EOF {
			if len(line) == 0 {
				return
			}
			fmt.Println(line, isprefix)
			return
		}
		if err != nil {
			fmt.Println("read line err", err)
			continue
		}
		//fmt.Println(line)
		count += 1
		if count % 2000000 == 0 {
			fmt.Println(string(line), isprefix)	
		}
	}
}

func durationTime(start int64, t string) int64 {
	end := time.Now().UnixNano()
	if t == "ms" {
		return (end - start) / int64(time.Millisecond)
	}
	return (end - start) / int64(time.Second)
}


func main() {
	file := os.Args[1]
	fmt.Println("file:", file)
	start := time.Now().UnixNano()
	ReadFlieBufio(file)
	fmt.Printf("Duration time %d ms.\n", durationTime(start, "ms"))
	//ReadFile(dir)
	// testIoutilReadFile(dir)
}
