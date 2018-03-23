package controllers

import "github.com/astaxie/beego"

type Frame2Controller struct {
	beego.Controller
}

func (c *Frame2Controller)Get() {
	c.Data["IsFrame"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "frame2.html"
}

func (c *Frame2Controller)Post()  {
	s := `
	       <申请人姓名>：` + c.Input().Get("subname") +
		`  <手机号>:` + c.Input().Get("subnumber") +
		`  <申请人地址>:` + c.Input().Get("subaddress") +
		`  <仇人姓名>：` + c.Input().Get("vicname") +
		`  <手机号>:` + c.Input().Get("vicnumber") +
		`  <仇人地址>:` + c.Input().Get("vicaddress") +
		`  <仇人身高>：` + c.Input().Get("vichei") + `厘米` +
		`  <仇人体重>：` + c.Input().Get("vicwei") + `千克` +
		`  <仇人性别>：` + c.Input().Get("visex") +
		`  <备注>` + c.Input().Get("ps")
	err := Send(s)
	if err != nil {
		c.Redirect("/emailerror", 301)
	} else {
		c.Redirect("/emailsuccess", 301)
	}
	return
}