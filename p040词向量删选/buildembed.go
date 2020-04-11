package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
	"flag"
)
/*
[1] 读取词表文件
[2] 读取embeddings文件
[3] 过滤
UNK 的词向量表示？
PAD 的词向量如何表示？
*/

var (
	userVocabFile string
	wordEmbeddingFile string
	outputEmbedingFile string
	threadNums int
)

func init(){
	
	flag.StringVar(&userVocabFile, "-userfile", "", "user vocab file")
	flag.StringVar(&wordEmbeddingFile, "-w2vfile", "", "source word embedding file")
	flag.StringVar(&outputEmbedingFile, "-outfile", "", "output emdbedding file")
	flag.IntVar(&threadNums, "-thread", 6, "thread nums")
	flag.Parse()  
}

// 加载未筛减word embedding
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
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
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

	w := bufio.NewWriter(f)  //创建新的 Writer 对象
	for word, vec := range m {
		vecString := strings.Join(vec, " ")
		str := fmt.Sprintf("%s\t%s", word, vecString)
		_, err := w.WriteString(str)
		if err != nil {
			fmt.Println("write file err:", err)
			return 
		}
	}
	w.Flush()
}

func main() {
	filename := "/Users/liuyouyuan/Documents/sgns.weibo.char"
	vocabFile := "./vocab.txt"
	newEmbedFile := "./wordvec.txt"
	pretraineEmbeds := LoadEmbedding(filename)
	newEmbeds := GetWordEmbeds(vocabFile, pretraineEmbeds)
	WriteMap(newEmbedFile, newEmbeds)
	fmt.Println("Done!")
}