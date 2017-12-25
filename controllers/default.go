package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	sess := c.StartSession()
	username := sess.Get("username")
	fmt.Println(username)
	c.TplName = "index.tpl"
}

func (c *MainController)Post(){
	var user models.User
	intputs := c.Input()
	user.Username = intputs.Get("username")
	user.Pwd = intputs.Get("pwd")
	fmt.Println(user)
}
