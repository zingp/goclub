package main

import (
	"strings"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

type ProcObj struct {
	Name string
	StartCmd string
	TimeInterval int
}

//AppConf 存放配置文件
type AppConf struct {
	ProcMaP map[string]*ProcObj
}

var appConf = &AppConf{
	ProcMaP: make(map[string]*ProcObj),
}

func initConfig(file string) (err error) {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	timeInterval := conf.DefaultInt("timeInterval", 5)
	procConfStr := strings.TrimSpace(conf.String("proc"))
	procConfSlice := strings.Split(procConfStr, ",")
	for _, perProc := range procConfSlice {
		if len(perProc) == 0 {
			continue
		}
		perProcStr := strings.TrimSpace(perProc)
		perProcSlice := strings.Split(perProcStr, ":")
		if len(perProcSlice) == 2 {
			name := strings.TrimSpace(perProcSlice[0])
			startCmd := strings.TrimSpace(perProcSlice[1])
			procObj := &ProcObj{
				Name: name,
				StartCmd: startCmd,
				TimeInterval: timeInterval,
			}
			_, ok := appConf.ProcMaP[name]
			if !ok {
				appConf.ProcMaP[name] = procObj
			}

			continue
		}
		fmt.Printf("monitor.cfg item:proc invalid")
		logs.Error("monitor.cfg item:proc invalid")
	}
	
	return
}