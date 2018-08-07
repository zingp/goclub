package main

import (
	"time"
	"github.com/astaxie/beego/logs"
)

// 超时或者日志上传完毕关闭管道，退出各个goroutine
func goroutineExit() {
	defer waitGroup.Done()
	defer close(hadpChan)
	defer close(lzopChan)
	defer close(rsyncChan)

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {

		if int(time.Now().Unix()-timeStart) > (appConf.timeout * 60) {
			logs.Error("timeout:", appConf.timeout)
			return
		}

		if sucessNum == int32(len(hostMap)) {
			logs.Info("all done.")
			return
		}
	}
}
