package main

import (
	"go_dev/s23_conf/config"
	"fmt"
)

func main() {
	confFile := "./test.cfg"
	conf, err := config.NewConfig(confFile)
	if err != nil {
		fmt.Printf("read config file err: %v", err)
		return
	}

	hostname, err := conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v", err)
		return
	}
	fmt.Println("Hostname:", hostname)

	port, err := conf.GetInt("port")
	if err != nil {
		fmt.Printf("get port err: %v", err)
		return
	}
	fmt.Println("Port:", port)
}

