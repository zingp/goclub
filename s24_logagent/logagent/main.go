package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

var appConfig = &AppConfig{}

func main() {
	err := initConfig("./logagent.cfg")
	if err != nil {
		fmt.Println("init config error:", err)
		return
	}
	fmt.Printf("Load config:\n%v\n", appConfig)

	err = initLogs()
	if err != nil {
		fmt.Println("init logs error:", err)
		return
	}

	err = initKafka()
	if err != nil {
		logs.Error("init kafka error:%v", err)
		return
	}

	RunServer()	
}
