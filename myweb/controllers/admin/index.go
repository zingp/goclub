package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["moduel"] = "admin"
	c.Data["page"] = "admin.Index"
	c.TplName = "admin/index.html"
}
