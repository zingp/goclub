/*
[1] 读取词表文件
[2] 读取embeddings文件
[3] 过滤
UNK 的词向量表示？
PAD 的词向量如何表示？
*/
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
)

func LoadVocab(path string) map[string]bool {
	vocab := map[string]bool{}
	fd, err := os.Open(path)
	if err != nil {
		fmt.Printf("open %s error.", path)
		return vocab
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if err == io.EOF {
			lineSlice := strings.Split(line, " ")
			if len(lineSlice) != 0 {
				word := lineSlice[0]
				vocab[word] = true
			}
			break
		}
		lineSlice := strings.Split(line, " ")
		if len(lineSlice) != 0 {
			word := lineSlice[0]
			vocab[word] = true
		}
	}
	return vocab
}

func LoadInEmbeds(path string, strCh chan string) {
	defer wg.Done()
	fd, err := os.Open(path)
	if err != nil {
		fmt.Printf("open %s error: %v", path, err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
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

func FilterAndWrite(path string, m map[string]bool, strCh chan string) {
	defer wg.Done()

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("create  %s err: %v", path, err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for line := range strCh {
		lineSlice := strings.Split(line, " ")
		if len(lineSlice) == 0 {
			continue
		}
		word := lineSlice[0]
		if _, ok := m[word]; ok {
			_, err = w.WriteString(line)
			if err != nil {
				fmt.Println("write file err:", err)
				continue
			}
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
	vocabFile  string
	inEmbFile  string
	outEmbFile string
	wg         = sync.WaitGroup{}
)

func init() {
	flag.StringVar(&vocabFile, "v", "", "user vocab file")
	flag.StringVar(&inEmbFile, "i", "", "source word embedding file")
	flag.StringVar(&outEmbFile, "o", "./outemb.txt", "output emdbedding file")
	flag.Parse()
}

func main() {
	strCh := make(chan string, 100)
	start := time.Now().UnixNano()
	vocabMap := LoadVocab(vocabFile)
	wg.Add(1)
	go LoadInEmbeds(inEmbFile, strCh)

	wg.Add(1)
	go FilterAndWrite(outEmbFile, vocabMap, strCh)

	wg.Wait()

	t := durationTime(start, "s")
	fmt.Printf("Write filter embedding to file: %s\n", outEmbFile)
	fmt.Printf("Duration %d s.\n", t)
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o embfilter buildembed.go
