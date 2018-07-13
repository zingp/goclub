package main
/*
建立与服务端的链接
进行数据收发
关闭链接
*/
import (
	"strings"
	"bufio"
	"net"
	"fmt"
	"os"
)

func SocketClient(){
	conn, err := net.Dial("tcp", "localhost:8800")
	if err!=nil {
		fmt.Println("Dial failed:", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input,_ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		if trimInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimInput))
		if err != nil {
			return
		}

		// 接收服务端数据
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		fmt.Println("Read Data:", string(buf))
	}
}

func main(){
	SocketClient()
}