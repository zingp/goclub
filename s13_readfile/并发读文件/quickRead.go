package main
import(
	"os"
	"io"
	"sync"
	"bufio"
	"fmt"
	"strings"
)
const mb = 1024 * 1024
const gb = 1024 * mb

func main() {
	wg := sync.WaitGroup{}

	//这个通道用于发送各种goroutine中的每个已读单词。
	channel := make(chan (string))

	// 存储唯一单词计数的字典。
	dict := make(map[string]int64)

	//done是一个通道，所有的单词都已输入字典后的信号。
	done := make(chan (bool), 1)

	// 读取通道中所有输入的单词并将它们添加到字典中。
	go func() {
		for s := range channel {
			dict[s]++
		}

		//向主线程发出信号，表明所有的单词都已进入词典。
		done <- true
	}()

	// current表示文件的字节数。
	var current int64
	// Limit表示每个线程要处理的文件块大小。
	var limit int64 = 500 * mb
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			read(current, limit, "gameofthrones.txt", channel)
			fmt.Printf("%d thread has been completed \n", i)
			wg.Done()
		}()
		// 将current值增加1+(前一个线程读取的最后一个字节)。
		current += limit + 1
	}
	// Wait for all go routines to complete.
	wg.Wait()
	close(channel)
	// Wait for dictionary to process all the words.
	<-done
	close(done)
}

func read(offset int64, limit int64, fileName string, channel chan (string)) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	// 将文件指针移动到指定块的开始位置。
	file.Seek(offset, 0)
	reader := bufio.NewReader(file)
	// 这段代码确保chunk的开头是一个新单词。
        //如果在给定的位置遇到一个字符，它将移动几个字节直到单词的末尾。
	if offset != 0 {
		_, err = reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}
		if err != nil {
			panic(err)
		}
	}
	var cummulativeSize int64
	for {
		// 如果读大小超过了块大小，则断开。
		if cummulativeSize > limit {
			break
		}
		b, err := reader.ReadBytes(' ')
		// 如果遇到文件结束，则中断。
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		cummulativeSize += int64(len(b))
		s := strings.TrimSpace(string(b))
		if s != "" {
			// 将通道中的已读单词发送到字典中。
			channel <- s
		}
	}
}