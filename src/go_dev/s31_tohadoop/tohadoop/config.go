package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"os"
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
	timeout   int
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
	appConf.timeout = conf.DefaultInt("timeout", 45)
	return
}

func initDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		cmd := fmt.Sprintf(`mkdir -p %s`, dir)
		_, errRet := ExecCmdLocal(cmd)
		if errRet != nil {
			logs.Error("make dir:%s failed:%v", dir, err)
			return
		}
	}
}

func initHadoopDir() {
	hdfs := fmt.Sprintf(`%s%s/%s`, appConf.hdfs, timeMon, timeDay)
	cmdFmt := `if ! %s fs -ls %s >/dev/null 2>&1;then %s fs -mkdir -p %s;fi`
	cmd := fmt.Sprintf(cmdFmt, appConf.hadoopClient, hdfs, appConf.hadoopClient, hdfs)
	_, err := ExecCmdLocal(cmd)
	if err != nil {
		logs.Error("init hadoop dir:%s failed:%v", hdfs, err)
		return
	}
}
