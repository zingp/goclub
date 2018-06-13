package main

import (
	_ "goweb/hello/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

