package main

import (
	"bytes"
	"os/exec"
	"strings"
)

/*
执行shell命令
*/

func runShell(c string) (sucOut string, errOut string) {
	cmd := exec.Command("/bin/bash", "-c", c)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errOut = strings.Trim(string(stderr.Bytes()), "\n")
		return
	}

	sucOut = strings.Trim(string(stdout.Bytes()), "\n")
	return
}
