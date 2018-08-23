package main

import (
	"time"
	"go_dev/s23_conf/reconf"
	"sync/atomic"
	"fmt"
)

type AppConfig struct {
	hostname string
	port int
	kafkaAddr string
	kafkaPort int
}

// reload()协程写 和 for循环的读，都是对Appconfig对象，因此有读写冲突
type AppConfigMgr struct {
	config atomic.Value
}

// 初始化结构体
var appConfigMgr = &AppConfigMgr{}

func (a *AppConfigMgr)Callback(conf *reconf.Config) {
	appConfig := &AppConfig{}
	hostname, err := conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v\n", err)
		return
	}
	appConfig.hostname = hostname

	kafkaPort, err := conf.GetInt("kafkaPort")
	if err != nil {
		fmt.Printf("get kafkaPort err: %v\n", err)
		return
	}
	appConfig.kafkaPort = kafkaPort

	appConfigMgr.config.Store(appConfig)

}

func initConfig(file string) {
	// [1] 打开配置文件
	conf, err := reconf.NewConfig(file)
	if err != nil {
		fmt.Printf("read config file err: %v\n", err)
		return
	}

	// 添加观察者
	conf.AddObserver(appConfigMgr)

	// [2]第一次读取配置文件
	var appConfig AppConfig
	appConfig.hostname, err = conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v\n", err)
		return
	}
	fmt.Println("Hostname:", appConfig.hostname)

	appConfig.kafkaPort, err = conf.GetInt("kafkaPort")
	if err != nil {
		fmt.Printf("get kafkaPort err: %v\n", err)
		return
	}
	fmt.Println("kafkaPort:", appConfig.kafkaPort)

	// [3] 把读取到的配置文件数据存储到atomic.Value
	appConfigMgr.config.Store(&appConfig)
	fmt.Println("first load sucess.")

}

func run(){
	for {
		appConfig := appConfigMgr.config.Load().(*AppConfig)

		fmt.Println("Hostname:", appConfig.hostname)
		fmt.Println("kafkaPort:", appConfig.kafkaPort)
		fmt.Printf("%v\n", "--------------------")
		time.Sleep(5 * time.Second)
	}
}

func main() { 
	confFile := "../parseConfig/test.cfg"
	initConfig(confFile)
	// 应用程序 很多配置已经不是存在文件中而是etcd
	run() 
}