package main

import (
	"fmt"
	"time"
)

func test(){
	time.Sleep(time.Millisecond * 100)  //毫秒
}

func main() {
	now := time.Now()
	fmt.Println("当前时间：", now)   //获取当前本地时间

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

        // 格式化
	fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
	fmt.Println(now.Format("2006-01-02 15:04：05"))   //also like this..

        start := time.Now().UnixNano()  // 纳秒
	test()
	end := time.Now().UnixNano()

	fmt.Printf("cost: %d us\n", (end -start)/1000)

}
