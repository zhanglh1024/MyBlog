package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)

type User struct {
	Id       int 
	Username string
	Pwd      string
}

func getLink() beedb.Model {
	db, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*myblog/root/root")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}



func SaveUser(user User) error {
	orm := getLink()
	fmt.Println(user)
	err := orm.Save(&user)
	return err
}

func ValidateUser(user User) error {
	orm := getLink()
	var myuser User
	orm.Where(" Username = ?", user.Username).Find(&myuser)
	if myuser.Username != user.Username || myuser.Pwd != user.Pwd {
		return errors.New("用户名或密码错误！")
	}
	return nil
}