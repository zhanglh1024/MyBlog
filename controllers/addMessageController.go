package controllers

import "github.com/astaxie/beego"

type AddMessageController struct {
	beego.Controller
}

func (c *AddMessageController)Get(){
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "message_add.html"
}