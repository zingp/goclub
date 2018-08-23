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

var etcdKey = "/logagent/10.136.10.191/logconfig"
var etcdValue = `[
	{
	"service":"test_service",
	"log_path": "/search/nginx/logs/ping-android.shouji.sogou.com_access_log","topic": "nginx_log",
	"send_rate": 1000
	},
	{
	"service":"srv.android.shouji.sogou.com",
	"log_path": "/search/nginx/logs/srv.android.shouji.sogou.com_access_log","topic": "nginx_log",
	"send_rate": 2000
	}
]`

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

	// put
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	ret, err := cli.Put(ctx, etcdKey, etcdValue)
	cancel()
	if err != nil {
		fmt.Println("put data to etcd failed, error:", err)
		return
	}
	fmt.Println("ret:", ret)

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, etcdKey)
	cancel()
	if err != nil {
		fmt.Println("get data from etcd failed, error:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("etcd key = %s , etcd value = %s", ev.Key, ev.Value)
	}
}

