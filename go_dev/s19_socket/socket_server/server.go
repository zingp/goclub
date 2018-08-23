package main
import(
	"fmt"
	"net"
)

/*
1. 监听端口
2.接收客户端的链接
3.创建goroutine,处理该连接
*/

func SocketServer() {
	fmt.Println("Start Server:")
	listen, err := net.Listen("tcp","0.0.0.0:8800")
	if err != nil {
		fmt.Println("failed to listen:",err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn){
	defer conn.Close()

	for{
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		fmt.Println("Read Data:", string(buf))

		//返回客户端数据
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("Send failed:", err)
			return
		}
	}
}

func main() {
	SocketServer()
}