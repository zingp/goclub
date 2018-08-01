package main

import (
	"github.com/astaxie/beego/logs"
	"os/exec"
	"strings"
)

// ExecCmdLocal is func run shell cmd on local host
func ExecCmdLocal(cmd string) (output string, err error) {
	res := exec.Command("/bin/sh", "-c", cmd)
	cont, err := res.Output()
	if err != nil {
		logs.Error("run shell cmd:%s, error:%v", cmd, err)
		return
	}

	output = strings.Trim(string(cont), "\n")
	return
}
