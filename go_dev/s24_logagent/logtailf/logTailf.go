package main

import(
	_ "github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"fmt"
	"time"
)


func testTail(filename string) {
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		Location:&tail.SeekInfo{ Offset:0, Whence: 2},  // 读到末尾
		MustExist: false,  //不存在也不报错
		Poll: true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok =<- tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 *time.Millisecond)
			continue
		}
		fmt.Println("LINE:", msg.Text)
	}
}

func main() {
	file := "./test.log"
	testTail(file)
}