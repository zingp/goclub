package main

import (
	"fmt"
	_ "myweb/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.AddTemplateExt("xxx")  // 自定义模板后缀

	beego.SetStaticPath("/code", "static")
	beego.BConfig.WebConfig.DirectoryIndex = true  //默认访问目录会列出文件

	// app.conf 中直接配mysqlhost = 10.143.57.61
	// host := beego.AppConfig.String("mysqlhost")
	// port, _ := beego.AppConfig.Int("mysqlport")
	// fmt.Println(host, port)

	/*
	app.conf 中配:
	[dbconfig]
	mysqlhost = 10.143.57.61
	msqlport = 3306
	*/
	host := beego.AppConfig.String("dbconfig::mysqlhost")
	port, _ := beego.AppConfig.Int("dbconfig::mysqlport")
	fmt.Println(host, port)

	beego.Run()
}

