package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})                 //登录页面
	beego.Router("/regist", &controllers.RegistController{})				//注册页面
}
