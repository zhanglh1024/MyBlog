package models

import (
	"time"
	"fmt"
)

type Message struct {
	Id			 int
	Uid        	 int     //留言者ID
	Content 	 string //`orm:"size(2000)"` //留言内容
	Attachment   string    //附件
	Created      time.Time //`orm:"index"` //留言时间
	Auther       string    //留言人姓名
	ReplayCount  int     //评论个数
	ReplayId     int     //评论用户ID
}

func GetAllMessage() ([]Message,error) {
	orm:=getLink()
	var message []Message
	err := orm.FindAll(&message)
	fmt.Println(err)
	return message,err
}

func InsertMessage(message Message) error {
	orm:=getLink()
	err:= orm.Save(&message)
	return err
}