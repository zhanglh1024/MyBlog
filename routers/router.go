package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})                 //登录页面
	beego.Router("/regist", &controllers.RegistController{})				//注册页面
	beego.Router("/message", &controllers.MessageController{})				//留言页面
	beego.Router("/add/message", &controllers.AddMessageController{})		//添加留言页面
	beego.Router("/police",&controllers.PoliceController{})				//警察页面
}
