package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"fmt"
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
	err := models.SaveUser(user)
	if err == nil{
		c.Redirect("/login",301)
	}else{
		fmt.Println(err)
		c.Redirect("/",301)
	}
	return
}