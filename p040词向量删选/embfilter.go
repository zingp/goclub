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
	"time"
	"sync"
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


func LoadInEmbeds(path string, strCh chan string){
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
			close(strCh)
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
			break
		}
		strCh <- line
	}
	
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

// 加载未筛减word embedding
func LoadEmbedding(f string) map[string][]string {
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
			if len(line) != 0 {
				sliceLine := strings.Split(line, " ")
				key := sliceLine[0]
				value := sliceLine[1:]
				embedings[key] = value
				break
			}
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
			break
		}
		// fmt.Println(line)
		if n == 0 {
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

// 传入用户词表文件，得到用户词表的wordembedings
func GetWordEmbeds(f string, m map[string][]string) map[string][]string {
	embedings := map[string][]string{}

	file, err := os.Open(f)
	if err != nil {
		fmt.Println("open file err:", err)
		return embedings
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			// 最后一行没有"\n"的情况
			sliceLine := strings.Split(line, " ")
			word := sliceLine[0]
			if vec, ok := m[word]; ok {
				embedings[word] = vec
			}
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		sliceLine := strings.Split(line, " ")
		word := sliceLine[0]
		// id := sliceLine[1]

		if vec, ok := m[word]; ok {
			embedings[word] = vec
		}
	}
	return embedings
}

func WriteMap(file string, m map[string][]string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	defer w.Flush()
	// 写入第一行
	firstLine := fmt.Sprintf("%d %d\n", len(m), 200)
	_, err = w.WriteString(firstLine)
	if err != nil {
		fmt.Println("write file err:", err)
		return
	}

	for word, vec := range m {
		vecString := strings.Join(vec, " ")
		str := fmt.Sprintf("%s %s", word, vecString)
		_, err = w.WriteString(str)
		if err != nil {
			fmt.Println("write file err:", err)
			break
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
	userVocabFile      string
	inEmbFile  string
	outEmbFile string
	wg       = sync.WaitGroup{}
)

func init() {
	flag.StringVar(&userVocabFile, "v", "", "user vocab file")
	flag.StringVar(&inEmbFile, "i", "", "source word embedding file")
	flag.StringVar(&outEmbFile, "o", "./outemb.txt", "output emdbedding file")
	flag.Parse()
}

func main() {
	strCh := make(chan string, 100)
	start := time.Now().UnixNano()
	vocabMap := LoadVocab(userVocabFile)
	wg.Add(1)
	go LoadInEmbeds(inEmbFile, strCh)

	wg.Add(1)
	go FilterAndWrite(outEmbFile, vocabMap, strCh)

    wg.Wait()
	// pretraineEmbeds := LoadEmbedding(wordEmbeddingFile)
	// end := time.Now().Second()
	// fmt.Printf("Load word embeddings cost %d s.\n", (end - start))
	// newEmbeds := GetWordEmbeds(userVocabFile, pretraineEmbeds)
	// WriteMap(outputEmbedingFile, newEmbeds)
	t := durationTime(start, "s")
	fmt.Printf("Write filter embedding to file: %s\n", outEmbFile)
	fmt.Printf("Duration %d s.\n", t)
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o embfilter buildembed.go