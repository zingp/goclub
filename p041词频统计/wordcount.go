package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	wordCountMap = make(map[string]int, 0)
	wg           = sync.WaitGroup{}

	a, h    bool
	inFile  string
	outFile string
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

func WordCount(strCh chan string) {
	defer wg.Done()

LOOP:
	for {
		select {
		case line, ok := <-strCh:
			if !ok {
				break LOOP
			}
			newLine := strings.TrimRight(line, "\n")
			newLine = strings.TrimSpace(newLine)
			sliceLine := strings.Split(newLine, " ")
			for _, word := range sliceLine {
				if len(word) == 0 {
					continue
				}
				wordCountMap[word]++
			}
		}
	}
	return
}

func WritePairListToFile(filePath string, p PairList) {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, pair := range p {
		line := fmt.Sprintf("%s %d\n", pair.Key, pair.Value)
		_, err = w.WriteString(line)
		if err != nil {
			fmt.Println("write file err:", err)
			return
		}
	}
	w.Flush()
}

func durationTime(start int64, t string) int64 {
	end := time.Now().UnixNano()
	if t == "ms" {
		return (end - start) / int64(time.Millisecond)
	}
	return (end - start) / int64(time.Second)
}

func init() {
	flag.BoolVar(&a, "a", false, "Ascending order by word frequency")
	flag.BoolVar(&h, "h", false, "Prints help information")
	flag.StringVar(&inFile, "i", "", "Input file")
	flag.StringVar(&outFile, "o", "./output.txt", "Output file")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, `wordcount version: wordcount/1.1.0
Usage: wordcount [-ha] [-i filename] [-o filename]

Options:
`)
	flag.PrintDefaults()
}

func main() {

	if h {
		flag.Usage()
		return
	}

	if inFile == "" {
		fmt.Println("Parameter -i is required.")
		flag.Usage()
		return
	}

	start := time.Now().UnixNano()
	strCh := make(chan string, 100)

	wg.Add(1)
	go ReadFile(inFile, strCh)

	wg.Add(1)
	go WordCount(strCh)

	wg.Wait()

	pairSlice := mapToSlice(wordCountMap)
	if a {
		sort.Sort(pairSlice)
	} else {
		sort.Sort(sort.Reverse(pairSlice))
	}

	WritePairListToFile(outFile, pairSlice)
	fmt.Printf("Output to file: %s.\n", outFile)
	fmt.Printf("Duration: %d s.\n", durationTime(start, "s"))
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wordcount *.go
