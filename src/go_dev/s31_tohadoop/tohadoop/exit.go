package main

import (
	"github.com/astaxie/beego/logs"
	"time"
)

// 超时或者日志上传完毕关闭管道，退出各个goroutine
func goroutineExit() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		func() {
			defer close(rsyncChan)
			defer close(lzopChan)
			defer close(hadpChan)

			if int(time.Now().Unix()-timeStart) > (appConf.timeout * 60) {
				logs.Error("timeout:", appConf.timeout)
				return
			}

			if sucessNum == int32(len(hostMap)) {
				logs.Info("all done.")
				return
			}
		}()
	}
}
