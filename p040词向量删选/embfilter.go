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


var (
	h          bool
	vocabFile  string
	inEmbFile  string
	outEmbFile string
	retSlice   = []string{}
	wg         = sync.WaitGroup{}
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

func Filter(m map[string]bool, strCh chan string) {
	defer wg.Done()

	retSlice = append(retSlice, "")

	count := 0
	dim := 0
	for line := range strCh {
		idx := strings.Index(line, " ")
		if idx < 0 {
			continue
		}
		word := line[:idx]
		if _, ok := m[word]; ok {
			retSlice = append(retSlice, line)
			count += 1
		}
	}

	if count > 1 {
		lineSlice := strings.Split(retSlice[1], " ")
		dim = len(lineSlice[1:])
	}

	retSlice[0] = fmt.Sprintf("%d %d\n", count, dim)
}

func WriteFile(s []string, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("create  %s err: %v", path, err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for _, line := range s {
		_, err = w.WriteString(line)
		if err != nil {
			fmt.Println("write file err:", err)
			continue
		}
	}
	return
}

// 暂时弃用
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
	go Filter(vocabMap, strCh)
	
	wg.Wait()
	WriteFile(retSlice, outEmbFile)

	t := durationTime(start, "s")
	fmt.Printf("Write to %s\n", outEmbFile)
	fmt.Printf("Duration %d s.\n", t)
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o embfilter embfilter.go
