package controllers

import "github.com/astaxie/beego"

type EmailSuccessController struct {
	beego.Controller
}

func (c *EmailSuccessController)Get()  {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "emailsuccess.html"
}