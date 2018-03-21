package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Owner"] = "张隆虎"
	c.Data["Email"] = "tianaishang@gmail.com"
	c.TplName = "index.tpl"
}


