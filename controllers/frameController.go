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
	user := "15659685590@sina.com" //默认发送邮件的账号密码
	password := "zlh984046"
	host := "smtp.sina.com:25"
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