package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
)

type RegistController struct {
	beego.Controller
}

func (c *RegistController)Get(){
	c.TplName = "regist.html"
}

func (c *RegistController)Post(){
	var user models.User
	intputs := c.Input()
	user.Username = intputs.Get("username")
	user.Pwd = intputs.Get("pwd")
	if models.SaveUser(user) == nil{
		c.Redirect("/login",301)
	}else{
		c.Redirect("/",301)
	}
	return
}