package main


import(
	"bufio"
	"os"
	"fmt"
	"math/rand"
	"time"
)


// 带缓冲从标准输入读
func testReadLine() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error info:", err)
		return
	}

	fmt.Println("您的输入是：", line)
}

// 只要实现Read 方法的结构体A，就可以用bufio.NewReader(A的实例)
type RandString struct {}

//随机字串 
func (r *RandString) Read(p []byte) (n int, err error) {
	str := "qwertyuiopasdfghjklzxcvbnm0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	for i:=0;i<32;i++ {
		index := rand.Intn(len(str))
		p[i] = str[index]
	}
	p[33] = '\n'

	return len(p), nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var randStr = &RandString{}
	randReader := bufio.NewReader(randStr)
	lineByte, prefix, _ := randReader.ReadLine()
	fmt.Printf("rand:%s prefix:%v \n", string(lineByte), prefix)
} 