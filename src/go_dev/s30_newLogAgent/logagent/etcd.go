package main

import (
	"context"
	"time"
	"fmt"

	client "github.com/coreos/etcd/clientv3"
)

var (
	etcdAddr = "10.134.123.183:2379"
	timeout time.Duration = 5
	etcdWatchKey = "/logagent/%s/logconfig"
	confChan = make(chan string)
)

func initEtcd() {

	cli, err := client.New(client.Config{
		Endpoints:   []string{etcdAddr},
		DialTimeout: timeout * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd error:", err)
		return
	}
	defer cli.Close()

	// 生成etcd key
	var etcdKey []string
	ips, err := getLocalIP()
	if err != nil {
		fmt.Println("get local ip error:", err)
		return 
	}
	for _, ip := range ips {
		key := fmt.Sprintf(etcdWatchKey, ip)
		etcdKey = append(etcdKey, key)
	}

	// 第一次运行主动从etcd拉取配置
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	for _, key := range etcdKey {
		resp, err := cli.Get(ctx, key)
		cancel()
		if err != nil {
			fmt.Println("get etcd key failed, error:", err)
			return
		}
		
		for _, ev := range resp.Kvs {
			// 返回的类型不是string
			// confChan <- stirng(ev.Value)
			fmt.Printf("etcd key = %s , etcd value = %s", ev.Key, ev.Value)
		}

		go etcdWatch(cli, key)
	}	
}


func etcdWatch(cli *client.Client, key string) {
	rch := cli.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}

//GetEtcdConfig is func get etcd conf
func GetEtcdConfig() chan string{
	return confChan
}