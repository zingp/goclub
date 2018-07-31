package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

var (
	etcdAddr     = []string{"10.134.123.183:2379"}
	etcdWatchKey = "/logagent/%s/logconfig"
)

func main() {
	// 读取本agent配置
	err := initLogs("./log/logagent.log", "debug")
	if err != nil {
		fmt.Printf("init log failed:%v", err)
		return
	}

	err = initEtcd(etcdAddr, etcdWatchKey, 5)
	if err != nil {
		logs.Error("initEtcd error:%v", err)
		return
	}

	runServer()
}
