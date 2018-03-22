package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"MyBlog/models"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get(){
	c.TplName = "login.html"
}

func (c *LoginController)Post(){
	var user models.User
	intputs := c.Input()
	user.Username = intputs.Get("username")
	user.Pwd = intputs.Get("pwd")
	autologin := c.Input().Get("autologin") == "on"
	fmt.Printf("user:%v",user,"---------------------")
	if models.ValidateUser(user) != nil{
		c.Redirect("/login",301)
		return
	}
	maxage := 0
	if autologin {
		maxage = 1 << 31 - 1
	}
	c.Ctx.SetCookie("username",user.Username,maxage,"/")
	c.Ctx.SetCookie("pwd",user.Pwd,maxage,"/")
	c.Redirect("/",301)
	return
}

func checkAccount(ctx *context.Context)bool{
	var user models.User
	ck,err := ctx.Request.Cookie("username")
	if err!=nil{
		return false
	}
	user.Username = ck.Value
	ck,err=ctx.Request.Cookie("pwd")
	if err!=nil{
		return false
	}
	user.Pwd = ck.Value
	if models.ValidateUser(user) == nil {
		return true
	}
	return false

}