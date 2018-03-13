package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main()  {

	conn, err := net.Dial("tcp", "localhost: 9999")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "q" {
			fmt.Println("God bye!")
			return
		}

		_, err = conn.Write([]byte(trimmedInput))

		if err !=  nil{
			return
		}
	}

}
