package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
)
func LoadEmbedding(f string) map[string][]string{
	embedings := map[string][]string{}
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("open file err:", err)
		return embedings
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	n := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
		}
		// fmt.Println(line)
		if n==0 {
			n += 1
			continue
		}
		sliceLine := strings.Split(line, " ")
		key := sliceLine[0]
		value := sliceLine[1:]
		embedings[key] = value
	}
	return embedings
}

func main() {
	filename := "/Users/liuyouyuan/Documents/sgns.weibo.char"
	embedings := LoadEmbedding(filename)
	fmt.Println(len(embedings))
}