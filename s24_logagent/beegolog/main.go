package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func testLogs() {
	config := make(map[string]interface{})
	config["filename"] = "./test.log"
	config["level"] = logs.LevelDebug
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)  // 打印文件名、文件行号

	logs.Debug("this is a Debug, my name is %s", "yy01")
	logs.Trace("this is a trace, my name is %s", "yy02")
	logs.Warn("this is a warn, my name is %s", "yy03")
}

func main() {
	testLogs()
}
