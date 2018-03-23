package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"net/smtp"
	"fmt"
)

type FrameController struct {
	beego.Controller
}

func (c *FrameController)Get()  {
	c.Data["IsFrame"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "frame.html"
}
func (c *FrameController) Post() {
	//获取表单中信息
	s := `
	       <申请人姓名>：` + c.Input().Get("subname") +
		`  <手机号>:` + c.Input().Get("subnumber") +
		`  <申请人地址>:` + c.Input().Get("subaddress") +
		`  <欠款人姓名>：` + c.Input().Get("vicname") +
		`  <手机号>:` + c.Input().Get("vicnumber") +
		`  <欠款人地址>:` + c.Input().Get("vicaddress") +
		`  <欠款额度>：` + c.Input().Get("num") +
		`  <备注>` + c.Input().Get("ps")
	err := Send(s) //调用邮件发送函数
	if err != nil {
		c.Redirect("/emailerror", 301) //如果发送失败跳转到失败提示页面
	} else {
		c.Redirect("/emailsuccess", 301) //如果发送成功跳转到成功提示页面
	}

	return
}
func SendToMail(user,password,host,sendto,subject,body,mailtype string)error{
	hp := strings.Split(host,":")
	auth := smtp.PlainAuth("",user,password,hp[0])
	var contentType string
	if mailtype == "html"{
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	}else {
		contentType = "Content-Type: text/" + "; charset=UTF-8"
	}


	msg := []byte("To: " + sendto + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + "\r\n" + contentType + "\r\n\r\n" + body)
	send_to := strings.Split(sendto, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)

	return err
}

func Send(s string) error  {
	user := "15659685590@163.com" //默认发送邮件的账号密码
	password := "zlh984046"
	host := "smtp.163.com:25"
	to := "2625041684@qq.com" //接收邮件的账号

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		` + s + `
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err) //在控制台输出错误信息
		//c.Redirect("/emailerror", 301)
	} else {
		fmt.Println("Send mail success!")
		//c.Redirect("/emailsuccess", 301)
	}
	return err
}