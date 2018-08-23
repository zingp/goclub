package main

import (
	"time"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"math/rand"
)

func initLogs(filename string) (err error) {
	config := make(map[string]interface{})
	config["filename"] = filename
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)
	return
}


func main() {
	rand.Seed(time.Now().UnixNano())
	err := initLogs("./log.log")
	if err != nil {
		fmt.Printf("init log error:%v", err)
		return
	}

	for {
		logs.Error("[Sogou-Observer,httpLib=1,a=1,b=2,c= 1,Owner=OP][400]")
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}