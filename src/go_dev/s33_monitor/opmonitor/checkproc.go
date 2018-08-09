package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

// 定时check 不存在则拉起
func checkProc(p string, start string, t time.Duration) {
	defer waitGroup.Done()

	ticker := time.NewTicker(t)
	cmdFmt := `ps aux |grep %s |grep -v grep`
	checkCmd := fmt.Sprintf(cmdFmt, p)
	for range ticker.C {
		// 检查进程是否存活
		sucOut, _ := runShell(checkCmd)
		logFmt := `proc [name=%s,cont=%s]`
		if len(sucOut) == 0 {
			runShell(start)
			s, e := runShell(checkCmd)
			if len(e) == 0 {
				logs.Info(fmt.Sprintf(logFmt, p, s))
			}
		}
	}
}
