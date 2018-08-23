package main

import "go_dev/s10_spider/httpReq"

func main() {

	url := "http://www.cnblogs.com/zingp/p/5878330.html"
	httpReq.GetHtml(url)
}