package main

import (
	// "time"
	"github.com/astaxie/beego/logs"
	"sync"
	"fmt"
)

var waitGroup sync.WaitGroup

func main() {
	logFile := "./monitor.log"
	err := initLogs(logFile)
	if err != nil {
		fmt.Printf("init log failed:%v\n", err)
		return
	}

	confFile := "./monitor.cfg"
	err = initConfig(confFile)
	if err != nil {
		logs.Error("init config failed:%v", err)
		return
	}

	// for k, v := range appConf.ProcMaP{
	// 	waitGroup.Add(1)
	// 	go checkProc(k, v.StartCmd, time.Duration(v.TimeInterval)* time.Second)
	// }

	err = countItems(appConf.CountFile)
	if err != nil {
		logs.Error("count items failed:%v", err)
	}
	fmt.Println(countItemsMap)

	waitGroup.Wait()
}