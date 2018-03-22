package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
)

type MessageController struct {
	beego.Controller
}

func (c *MessageController)Get(){
	c.Data["IsMessage"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	message,err:= models.GetAllMessage()
	if err !=nil {
		beego.Error(err)
	}else {
		c.Data["Messages"] = message
	}
	c.TplName = "message.html"
}
