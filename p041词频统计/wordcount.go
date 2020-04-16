package main

import (
	"os"
	"fmt"
	"time"
	"io"
	"flag"
	"bufio"
	"strings"
	"sync"
)

var (
	wordCountMap = make(map[string]int, 0)
	wg = sync.WaitGroup{}
	wordSyncMap sync.Map
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

	LOOP: for {
		select {
		case line, ok :=<- strCh:
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

func WordSyncCount(strCh chan string) {
	defer wg.Done()

	for {
		select {
		case line, ok :=<- strCh:
			if !ok {
				return 
			}
			newLine := strings.TrimRight(line, "\n")
			newLine = strings.TrimSpace(newLine)
			sliceLine := strings.Split(newLine, " ")
			for _, word := range sliceLine {
				if len(word) == 0 {
					continue
				}
				// 添加
				v, ok := wordSyncMap.Load(word)
				if ok {
					if value, ok := v.(int); ok {
						wordSyncMap.Store(word, value + 1)
					}
				} else {
					wordSyncMap.Store(word, 1)
				}
			}
		}
	}
}


func WritePairListToFile(filePath string, p PairList) {
	f, err := os.Create(filePath) 
	if err != nil {
		fmt.Println("open file err:", err)
		return 
	}
	defer f.Close()

	//创建新的 Writer 对象
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
		return (end - start)/int64(time.Millisecond)
	}
	return (end - start)/int64(time.Second)
}

var (
	inFile string
	outFile string
	threads int
)
	
func init(){
	flag.StringVar(&inFile, "i", "", "input file")
	flag.StringVar(&outFile, "o", "", "output file")
	flag.IntVar(&threads, "t", 4, "thread nums")
	flag.Parse()  
}

func main() {
	start := time.Now().UnixNano()
	strCh := make(chan string, 100)

	wg.Add(1)
	go ReadFile(inFile, strCh)

	for i:=0; i<=threads; i++ {
		wg.Add(1)
		// go WordCount(strCh)
		go WordSyncCount(strCh)
	}
	//wg.Add(1)
	// go WordCount(strCh)

	wg.Wait()

	//sortSlice := sortMapByValue(wordCountMap)
	sortSlice := sortSyncMapByValue(wordSyncMap)
	WritePairListToFile(outFile, sortSlice)
	
	fmt.Printf("Cost time %d s.\n", durationTime(start, "s"))
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wordcount wordcount.go