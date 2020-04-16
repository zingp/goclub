package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	gojieba "github.com/yanyiwu/gojieba"
)

func ReadFile(filePath string, strCh chan string) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			strCh <- line
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
			break
		}
		strCh <- line
	}
	close(strCh)

	return
}

// 这里不能起多个
func WordCut(strCh chan string, cutCh chan string, m bool, fc func(string, bool) []string) {
	defer wg.Done()

	for {
		select {
		case line, ok := <-strCh:
			if !ok {
				close(cutCh)
				return
			}
			newLine := strings.TrimRight(line, "\n")
			newLine = strings.TrimSpace(newLine)
			wordSlice := fc(newLine, m)
			cutLine := strings.Join(wordSlice, " ")
			cutCh <- cutLine
		}
	}
}

func WriteFile(filePath string, cutCh chan string) {
	defer wg.Done()

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for line := range cutCh {
		_, err = w.WriteString(line + "\n")
		if err != nil {
			fmt.Println("write file err:", err)
			return
		}
	}
	return
}

func durationTime(start int64, t string) int64 {
	end := time.Now().UnixNano()
	if t == "ms" {
		return (end - start) / int64(time.Millisecond)
	}
	return (end - start) / int64(time.Second)
}

var (
	wg      = sync.WaitGroup{}
	inFile  string
	outFile string
)

func init() {
	flag.StringVar(&inFile, "i", "", "input file")
	flag.StringVar(&outFile, "o", "", "output file")
	//flag.StringVar(&cutMode, "m", "cut", "cut mode: cut, search")
	flag.Parse()
}

func main() {
	start := time.Now().UnixNano()

	strCh := make(chan string, 10)
	cutCh := make(chan string, 10)

	jieba := gojieba.NewJieba()
	defer jieba.Free()

	wg.Add(1)
	go ReadFile(inFile, strCh)

	wg.Add(1)
	go WordCut(strCh, cutCh, true, jieba.Cut)

	wg.Add(1)
	go WriteFile(outFile, cutCh)

	wg.Wait()

	fmt.Printf("Cost time %d s.\n", durationTime(start, "s"))
}
