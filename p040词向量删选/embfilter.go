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
			idx := strings.Index(line, " ")
			if idx > 0 {
				word := line[:idx]
				vocab[word] = true
			}
			break
		}
		idx := strings.Index(line, " ")
		if idx > 0 {
			word := line[:idx]
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
		idx := strings.Index(line, " ")
		if idx < 0 {
			continue
		}
		word := line[:idx]
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

func usage() {
	fmt.Fprintf(os.Stderr, `embfilter version: embfilter/1.1.0
Usage: embfilter [-h] [-v filename] [-i filename] [-o filename]

Options:
`)
	flag.PrintDefaults()
}

var (
	h          bool
	vocabFile  string
	inEmbFile  string
	outEmbFile string
	wg         = sync.WaitGroup{}
)

func init() {
	flag.BoolVar(&h, "h", false, "Prints help information")
	flag.StringVar(&vocabFile, "v", "", "Vocab file")
	flag.StringVar(&inEmbFile, "i", "", "Source word embedding file")
	flag.StringVar(&outEmbFile, "o", "./outemb.txt", "Output emdbedding file")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	if h {
		flag.Usage()
		return
	}

	if vocabFile == "" || inEmbFile == "" {
		fmt.Println("Parameter -v -i  is required.")
		flag.Usage()
		return
	}

	strCh := make(chan string, 100)
	start := time.Now().UnixNano()
	vocabMap := LoadVocab(vocabFile)
	wg.Add(1)
	go LoadInEmbeds(inEmbFile, strCh)

	wg.Add(1)
	go FilterAndWrite(outEmbFile, vocabMap, strCh)

	wg.Wait()

	t := durationTime(start, "s")
	fmt.Printf("Write to %s\n", outEmbFile)
	fmt.Printf("Duration %d s.\n", t)
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o embfilter embfilter.go
