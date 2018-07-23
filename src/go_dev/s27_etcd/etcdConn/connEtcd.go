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

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	ret, err := cli.Put(ctx, "/logagent/conf", "logagent_value666")
	cancel()
	if err != nil {
		fmt.Println("put failed, error:", err)
		return
	}
	fmt.Println("ret:", ret)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed, error:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("etcd key = %s , etcd value = %s", ev.Key, ev.Value)
	}
}
