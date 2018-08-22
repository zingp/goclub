package main

import (
	"context"
	"fmt"
	client "github.com/coreos/etcd/clientv3"
	"time"
)

/*
启动etcd:
./etcd -listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2380
*/

func main() {
	// 连接etcd
	cli, err := client.New(client.Config{
		Endpoints:   []string{"10.134.123.183:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd error:", err)
		return
	}
	fmt.Println("connect success.")
	defer cli.Close()

	// 返回一个管道
	rch := cli.Watch(context.Background(), "/logagent/conf333")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
