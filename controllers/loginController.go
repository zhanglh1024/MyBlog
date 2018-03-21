package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"MyBlog/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get(){
	c.TplName = "login.tpl"
}

func (c *LoginController)Post(){
	var user models.User
	intputs := c.Input()
	user.Username = intputs.Get("username")
	user.Pwd = intputs.Get("pwd")
	fmt.Println(user,"---------------------")
	c.Redirect("/",301)
	return
}