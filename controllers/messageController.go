package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"time"
	"fmt"
)

type MessageController struct {
	beego.Controller
}

func (c *MessageController)Get(){
	c.Data["IsMessage"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	message,err:= models.GetAllMessage()
	if err !=nil {
		fmt.Println("留言跳转出错",err)
		beego.Error(err)
	}else {
		c.Data["Messages"] = message
	}
	c.TplName = "message.html"
}

func (c *MessageController)Add(){
	c.TplName = "message_add.html"
}

func (c *MessageController)Post(){
	if checkAccount(c.Ctx)==false {
		c.Redirect("/login",302)
		return
	}
	var message models.Message
	message.Auther = c.Input().Get("name")
	message.Content = c.Input().Get("message")
	message.Created = time.Now()

	err := models.InsertMessage(message)
	if err != nil{
		fmt.Println(err)
		beego.Error(err)
	}
	c.Redirect("/message",302)
	return
}
