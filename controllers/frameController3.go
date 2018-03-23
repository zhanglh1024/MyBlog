package controllers

import "github.com/astaxie/beego"

type Frame3Controller struct {
	beego.Controller
}

func (c *Frame3Controller)Get(){
	c.Data["IsFrame"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "frame3.html"
}

func (c *Frame3Controller)Post(){
	s := `
	       <申请人姓名>：` + c.Input().Get("subname") +
		`  <手机号>:` + c.Input().Get("subnumber") +
		`  <申请人地址>:` + c.Input().Get("subaddress") +
		`  <受害人姓名>：` + c.Input().Get("vicname") +
		`  <手机号>:` + c.Input().Get("vicnumber") +
		`  <受害人地址>:` + c.Input().Get("vicaddress") +
		`  <泼粪重量>：` + c.Input().Get("num") + `千克` +
		`  <备注>` + c.Input().Get("ps")
	err := Send(s)
	if err != nil {
		c.Redirect("/emailerror", 301)
	} else {
		c.Redirect("/emailsuccess", 301)
	}
	return
}
