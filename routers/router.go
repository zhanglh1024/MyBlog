package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/frame", &controllers.FrameController{})
	beego.Router("/login", &controllers.LoginController{})                 //登录页面
	beego.Router("/regist", &controllers.RegistController{})				//注册页面
	beego.Router("/message", &controllers.MessageController{})				//留言页面
	beego.Router("/add/message", &controllers.AddMessageController{})		//添加留言页面
	beego.Router("/police",&controllers.PoliceController{})				//警察页面
	beego.Router("/emailsuccess", &controllers.EmailSuccessController{})	//邮件发送成功
	beego.Router("/emailerror", &controllers.EmailErrorController{})		//邮件发送失败
	beego.Router("/frame2", &controllers.Frame2Controller{})
	beego.Router("/frame3", &controllers.Frame3Controller{})
	beego.Router("/frame4", &controllers.Frame4Controller{})
	beego.Router("/frame5", &controllers.Frame5Controller{})
	beego.Router("/addartical", &controllers.AddArticalController{})
}
