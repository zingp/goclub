package main

import(
	"os"
	"bufio"
	"io"
	"fmt"
)

// io.Copy实现cat命令
func Cat1(file string){
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()
	
	io.Copy(os.Stdout, f)
}


// 方法二：用bufio与Fprintf实现
func Cat(file string){
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		fmt.Fprintf(os.Stdout, "%s", data)
	}

}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("请指定文件名")
		return
	}
	for i:=1;i<len(os.Args);i++ {
		Cat1(os.Args[i])
	}
}