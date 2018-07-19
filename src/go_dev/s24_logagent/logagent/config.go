package main

import(
	"fmt"
	"strings"
	"github.com/astaxie/beego/config"
)

type AppConfig struct {
	ListenFile []string
	ThreadNum int
	KafkaAddr string
	KafkaTopic string
	LogFile string
	LogLevel string
}

func initConfig(file string) (err error){
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	listen_file := conf.String("listen_file")
	fileSlice := strings.Split(listen_file, ",")
	for _, item := range fileSlice {
		filename := strings.TrimSpace(item)
		if len(filename) == 0 {
			continue
		}
		appConfig.ListenFile = append(appConfig.ListenFile, filename)
	}
	

	appConfig.ThreadNum = conf.DefaultInt("DefaultInt", 8)
	appConfig.KafkaAddr = conf.String("kafka::addr")
	appConfig.KafkaTopic = conf.String("kafka::topic")
	appConfig.LogFile = conf.String("log::file")
	appConfig.LogLevel = conf.String("log::level")

	fmt.Println("ListenFile:", appConfig.ListenFile)
	fmt.Println("ThreadNum:", appConfig.ThreadNum)
	fmt.Println(":", appConfig.KafkaAddr)
	fmt.Println(":", appConfig.KafkaTopic)
	fmt.Println(":", appConfig.LogFile)
	fmt.Println(":", appConfig.LogLevel)
	

	return
}

