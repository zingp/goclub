package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"os"
	"strings"
)


func run_shell(s string) {
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
	run_shell(cmd)
}
