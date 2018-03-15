package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"os"
	"strings"
)

/*
执行shell命令
*/

func runShell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func main() {
	cmd := strings.Join(os.Args[1:], " ")
	fmt.Println("cmd>>>", cmd)
	runShell(cmd)
}
