package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

//AppConf 存放配置文件
type AppConf struct {
	remoteAddr   string
	localAddr    string
	lzopAddr     string
	logName      string
	hadoopClient string
	hdfs         string

	logFile   string
	hostFile  string
	threadNum int
}

func initConfig(file string) (err error) {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	appConf.remoteAddr = conf.String("remoteAddr")
	appConf.localAddr = conf.String("localAddr")
	appConf.lzopAddr = conf.String("lzopAddr")
	appConf.logName = conf.String("logName")
	appConf.hadoopClient = conf.String("hadoopClient")
	appConf.hdfs = conf.String("hdfs")
	appConf.logFile = conf.String("logFile")
	appConf.hostFile = conf.String("hostFile")
	appConf.threadNum = conf.DefaultInt("threadNum", 4)

	return
}
