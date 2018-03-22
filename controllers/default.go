package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Owner"] = "张隆虎"
	c.Data["Email"] = "tianaishang@gmail.com"
	c.Data["IsHome"] = true		//控制菜单上首页的高亮效果
	c.Data["IsLogin"] = checkAccount(c.Ctx) //调用函数判断账号密码时候正确来决定显示登录按钮还是退出按钮
	c.TplName = "index.html"
}


