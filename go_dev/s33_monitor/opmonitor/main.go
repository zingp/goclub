package main

import (
	"time"
	"fmt"
	"sync"
	"github.com/astaxie/beego/logs"
)

var waitGroup sync.WaitGroup

func main() {
	confFile := "./conf/app.cfg"
	err := initConfig(confFile)
	if err != nil {
		logs.Error("init config failed:%v", err)
		return
	}

	logFile := "./logs/monitor.log"
	err = initLogs(logFile)
	if err != nil {
		fmt.Printf("init log failed:%v\n", err)
		return
	}

	for k, v := range appConf.ProcMaP{
		waitGroup.Add(1)
		go checkProc(k, v.StartCmd, time.Duration(v.TimeInterval)* time.Second)
	}

	waitGroup.Add(1)
	go cron(appConf.CountFile)

	waitGroup.Wait()
}
