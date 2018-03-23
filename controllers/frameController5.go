package controllers

import "github.com/astaxie/beego"

type Frame5Controller struct {
	beego.Controller
}

func (c *Frame5Controller)Get()  {
	c.Data["IsFrame"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "frame5.html"
}

func (c *Frame5Controller)Post()  {
	s := `
	       <申请人姓名>：` + c.Input().Get("subname") +
		`  <手机号>:` + c.Input().Get("subnumber") +
		`  <申请人地址>:` + c.Input().Get("subaddress") +
		`  <受害人姓名>：` + c.Input().Get("vicname") +
		`  <手机号>:` + c.Input().Get("vicnumber") +
		`  <受害人地址>:` + c.Input().Get("vicaddress") +
		`  <受害人年龄>：` + c.Input().Get("num") + `岁` +
		`  <绑架地点>：` + c.Input().Get("vicplace") +
		`  <备注>` + c.Input().Get("ps")
	err := Send(s)
	if err != nil {
		c.Redirect("/emailerror", 301)
	} else {
		c.Redirect("/emailsuccess", 301)
	}
	return
}
