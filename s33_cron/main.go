package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func printLog(){
	log.Println("test cron")
}

func main(){
	c := cron.New()
	c.AddFunc("1 * * * *", printLog)
	c.Start()
	fmt.Println("start...")
	select{}
}