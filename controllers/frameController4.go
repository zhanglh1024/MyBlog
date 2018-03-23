package controllers

import (
	"github.com/astaxie/beego"
)

type Frame4Controller struct {
	beego.Controller
}

func (c *Frame4Controller)Get(){
	c.Data["IsFrame"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "frame4.html"
}

func (c *Frame4Controller)Post()  {
	s := `
	       <申请人姓名>：` + c.Input().Get("subname") +
		`  <手机号>:` + c.Input().Get("subnumber") +
		`  <申请人地址>:` + c.Input().Get("subaddress") +
		`  <指定保镖>：` + c.Input().Get("vicname") +
		`  <保护时间>:` + c.Input().Get("vicnumber") +
		`  <生活环境>:` + c.Input().Get("vicaddress") +
		`  <保镖数量>：` + c.Input().Get("num") +
		`  <备注>` + c.Input().Get("ps")
	err := Send(s)
	if err != nil {
		c.Redirect("/emailerror", 301)
	} else {
		c.Redirect("/emailsuccess", 301)
	}
	return
}
