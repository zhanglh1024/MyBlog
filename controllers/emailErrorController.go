package controllers

import "github.com/astaxie/beego"

type EmailErrorController struct {
	beego.Controller
}

func (c *EmailErrorController)Get()  {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "emailerror.html"
}