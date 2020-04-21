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
	h       bool
	inFile  string
	outFile string
	wg      = sync.WaitGroup{}
)

// ReadFile for read file
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

// CleanAndWrite for clean text and wirte to file
func CleanAndWrite(filePath string, strCh chan string) {
	defer wg.Done()

	fd, err := os.Create(filePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer fd.Close()

	w := bufio.NewWriter(fd)
	defer w.Flush()
	for line := range strCh {
		newLine := strings.ToLower(line)
		_, err := w.WriteString(newLine)
		if err != nil {
			fmt.Println("write file err:", err)
			continue
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
	fmt.Fprintf(os.Stderr, `tclean version: wordcount/1.1.0
Usage: tclean [-h] [-i filename] [-o filename]

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&h, "h", false, "Prints help information")
	flag.StringVar(&inFile, "i", "", "Input file")
	flag.StringVar(&outFile, "o", "./outclean.txt", "Output file")
	flag.Usage = usage
	flag.Parse()
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
	go CleanAndWrite(outFile, strCh)

	wg.Wait()
	t := durationTime(start, "s")
	fmt.Printf("Output to %s\n", outFile)
	fmt.Printf("Duration: %d s.\n", t)

	return
}
