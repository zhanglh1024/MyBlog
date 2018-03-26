package controllers

import (
	"github.com/astaxie/beego"
	"MyBlog/models"
	"time"
	"fmt"
)

type AddArticalController struct {
	beego.Controller
} 

func (c *AddArticalController)Get(){
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "addartical.html"
}

func (c *AddArticalController)Post()  {
	if checkAccount(c.Ctx)==false {
		return
	}
	var artical models.Artical
	artical.Title = c.Input().Get("title")
	artical.Auther = c.Input().Get("auther")
	artical.Content = c.Input().Get("content")
	artical.WritDate = time.Now()
	fmt.Println(artical,"Title:",artical.Title,"Auther:",artical.Auther)
	err := models.InsertArtical(&artical)
	if err != nil{
		beego.Error(err)
	}
	c.Redirect("/",302)
	return

}