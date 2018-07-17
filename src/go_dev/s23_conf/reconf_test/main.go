package main

import (
	"time"
	"go_dev/s23_conf/reconf"
	"fmt"
)

type AppConfig struct {
	hostname string
	port int
	kafkaAddr string
	kafkaPort int
}


func (a *AppConfig)Callback(conf *reconf.Config) {

	hostname, err := conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v\n", err)
		return
	}
	a.hostname = hostname
	

	kafkaPort, err := conf.GetInt("kafkaPort")
	if err != nil {
		fmt.Printf("get kafkaPort err: %v\n", err)
		return
	}
	a.kafkaPort = kafkaPort

}

var appConfig AppConfig

func main() {
	confFile := "../parseConfig/test.cfg"
	conf, err := reconf.NewConfig(confFile)
	if err != nil {
		fmt.Printf("read config file err: %v\n", err)
		return
	}

	conf.AddObserver(&appConfig)

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

	for {
		fmt.Println("Hostname:", appConfig.hostname)
		fmt.Println("kafkaPort:", appConfig.kafkaPort)
		fmt.Printf("%v\n", "--------------------")
		time.Sleep(5 * time.Second)
	}
}