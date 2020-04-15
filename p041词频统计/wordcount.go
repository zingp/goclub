package main

import (
	"os"
	"fmt"
	"time"
	//"runtime"
	"io"
	"flag"
	"bufio"
	"strings"
	"sync"
	"sort"
)
var wordCountMap = make(map[string]int, 0)
var wg = sync.WaitGroup{}

/* Map 排序 */
//要对golang map按照value进行排序，思路是直接不用map。
//用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序。
// A data structure to hold a key/value pair.
type Pair struct {
    Key   string
    Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p PairList) Len() int           { return len(p) }

func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
    p := make(PairList, len(m))
    i := 0
    for k, v := range m {
		p[i] = Pair{k, v}
		i += 1
	}
	// 降序
	sort.Sort(sort.Reverse(p))
	// sort.Sort(p)  // 升序
    return p
}


func ReadFile(filePath string, strCh chan string) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return 
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	strCh <- scanner.Text()
	// }
	
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
				fmt.Println("close chan")
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
	start := time.Now()
	strCh := make(chan string, 100)

	wg.Add(1)
	go ReadFile(inFile, strCh)
	wg.Add(1)
	go WordCount(strCh)
	wg.Wait()

	sortSlice := sortMapByValue(wordCountMap)
	
	WritePairListToFile(outFile, sortSlice)
	elapsed := time.Since(start)
	fmt.Printf("Cost time %d s.\n", elapsed/1e6)
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wordcount wordcount.go