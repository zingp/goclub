package main

import (
	"time"
	"io"
	"bufio"
	"os"
	"os/exec"
	"fmt"
	"strings"

	"github.com/astaxie/beego/config"
)

type AppConf struct {
	cmd string
	remoteAddr string
	localAddr string
	logName string
	hadoopClient string
	hdfs string
}
// statInfoLog.log.2018-07-31_13
// logfile: statInfoLog.log.2018-07-31_10.resin21.shouji.tc.ted.10.143.41.179.statInfoLog.log.lzo

var appConf AppConf
var hostMap map[string]string
var rsyncCmdChan = make(chan string, 60)

func initConfig(file string) (err error) {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	// appConf.cmd =  conf.String("cmd")
	appConf.remoteAddr = conf.String("remoteAddr")
	appConf.localAddr = conf.String("localAddr")
	appConf.logName = conf.String("logName")
	appConf.hadoopClient = conf.String("hadoopClient")
	appConf.hdfs = conf.String("hdfs")

	return
}

func getHostMap(file string)(err error){
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("open file:%s error:%v", file, err)
		return
	}
	defer f.Close()

	hostMap = make(map[string]string, 60)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			line = strings.TrimSpace(line)
			lineSlice := strings.Split(line,":")
			hostMap[lineSlice[0]] = lineSlice[1]
			break
		}
		if err != nil {
			continue
		}
		line = strings.TrimSpace(line)
		lineSlice := strings.Split(line,":")
		hostMap[lineSlice[0]] = lineSlice[1]
	}

	return
}

// statInfoLog.log   statInfoLog.log.2018-07-31_13
// remoteAddr= /search/odin/resin/WebContent/WEB-INF/log/statInfoLog/
// localAddr= /search/odin/resin/WebContent/WEB-INF/log/statInfoLog/
// 应该配置成本地
func genRsyncCmd(remoteAddr string, localAddr string, logName string){

	// 获取一小时前日志的名称
	h, _ := time.ParseDuration("-1h")
	lastHourTime := time.Now().Add(h).Format("2006-01-02_15")
	lastHourLog := fmt.Sprintf("%s.%s", logName, lastHourTime)

	// rsync 命令格式
	rsyncCmd := `rsync -avzP rsync.%s::odin%s%s %s`
	
	// 生成具体命令加入管道
	for key := range hostMap {
		cmd := fmt.Sprintf(rsyncCmd, key, remoteAddr, lastHourLog, localAddr)
		fmt.Println(cmd)
		rsyncCmdChan <- cmd
	}

}

func lzopLog(log string, zlog string) {
	//
}

func execCmdLocal(cmd string)(output string, err error){
	res := exec.Command("/bin/sh", "-c", cmd)
	cont, err := res.Output();
    if err != nil {
        fmt.Printf("run shell cmd:%s, error:%v", cmd, err)
        return
    }
    
	output = strings.Trim(string(cont), "\n")
	return
}

func main() {
	confFile := "./tohadoop.cfg"
	err := initConfig(confFile)
	if err != nil {
		fmt.Printf("load conf failed:%v\n", err)
	}

	hostFile := "./hosts"
	err = getHostMap(hostFile)
	if err != nil {
		fmt.Printf("get host map failed:%v\n", err)
	}

	genRsyncCmd(appConf.remoteAddr, appConf.localAddr, appConf.logName)
	
	// output, err:= execCmdLocal(appConfig.cmd)
	// if err != nil {
	// 	fmt.Printf("execCmdLocal error:%v",err )
	// }
	// fmt.Println(output)
}

// 判断命令执行成功