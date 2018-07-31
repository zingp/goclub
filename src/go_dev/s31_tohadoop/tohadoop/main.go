package main

import (
	"io"
	"bufio"
	"os"
	"os/exec"
	"fmt"
	"strings"

	"github.com/astaxie/beego/config"
)

type AppConfig struct {
	cmd string
}
// statInfoLog.log.2018-07-31_13
// logfile: statInfoLog.log.2018-07-31_10.resin21.shouji.tc.ted.10.143.41.179.statInfoLog.log.lzo

var appConfig AppConfig

func initConfig(file string) (err error) {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	// appConfig.cmd =  conf.String("cmd")
	hosts := conf.String("real_hosts")
	fmt.Println(hosts)

	return
}

func getHosts(file string){
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file:%d error:%v", file, err)
		return
	}
	defer f.Close()

	hostMap := make(map[string]string, 60)
	reader := bufio.NewReader(f)
	for {
		line, err:= reader.ReadString('\n')
		if err == io.EOF {
			line = strings.TrimSpace(line)
			lineSlice := strings.Split(line,":")
			hostMap[lineSlice[0]] = lineSlice[1]
			break
		}
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		lineSlice := strings.Split(line,":")
		hostMap[lineSlice[0]] = lineSlice[1]
	}
	fmt.Println(hostMap)
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
	getHosts(hostFile)

	// output, err:= execCmdLocal(appConfig.cmd)
	// if err != nil {
	// 	fmt.Printf("execCmdLocal error:%v",err )
	// }
	// fmt.Println(output)
}