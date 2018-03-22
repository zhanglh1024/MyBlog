package controllers

import "github.com/astaxie/beego"

type PoliceController struct {
	beego.Controller
}

func (c *PoliceController)Get()  {
	c.Data["IsPolice"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "police.html"
}