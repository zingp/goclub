package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Data["hello"] = "Hello World!"
	c.TplName = "my/hello.tpl"
}
