package main

import (
	"fmt"
	"math/rand"
)

// [1]定义负载均衡接口，及方法
type LoadBalance interface {
	DoBalance ([]string) string
}

// [2] 定义随机算法结构体 RandBalance，并实现LoadBalance接口
type RandBalance struct {}

func (r *RandBalance) DoBalance(addrList []string) string {
	// 此处省略判断ip是否是活的等业务逻辑
	l := len(addrList)
	index := rand.Intn(l)
	return addrList[index]
}

// [3] 定义轮询算法结构体 PollBalance，并实现LoadBalance接口
type PollBalance struct {
	curIndex int
}

func (p *PollBalance) DoBalance(addrList []string) string {
	// 此处省略判断ip是否是活的等业务逻辑
	index := p.curIndex % (len(addrList))   // 取余，防止index溢出
	addr := addrList[index]
	p.curIndex++

	return addr
}

// [4] 通过接口实现多态
func doBalance(balance LoadBalance, addrList []string) (addr string) {
	return balance.DoBalance(addrList)
}

func main() {
	// 随机生成IP地址切片
	var addrList []string
	for i:=0; i<5; i++ {
		addr := fmt.Sprintf("%d.%d.%d.%d:80", 
			rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
		addrList = append(addrList, addr)
	}
	fmt.Println("addrList:", addrList)

	// 通过配置balanceFunc变量，调用不同的算法
	var balanceFunc string = "poll"
	var balance LoadBalance
	if balanceFunc == "random" {
		balance = &RandBalance{}
	} else if balanceFunc == "poll" {
		balance = &PollBalance{}
	}

	// 10次调用负载均衡算法算出IP
	for i:=0; i<10; i++ {
		addr := doBalance(balance, addrList)
		fmt.Printf("Index %d  <--> Addr %s\n", i, addr)
	}
	
}