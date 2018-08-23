package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"os"
	"strings"
)

/*
执行shell命令
*/

func runShell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error detail:%s error status:%v\n",string(stderr.Bytes()), err)
		return
	}
	outStr := string(stdout.Bytes())
	fmt.Printf("Success out:%s\n", outStr)
}

func main() {
	cmd := strings.Join(os.Args[1:], " ")
	fmt.Println("cmd>>>", cmd)
	runShell(cmd)
}
