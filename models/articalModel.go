package models

import "time"

type Artical struct {
	Id 			int
	Title		string
	Auther		string
	Content     string
	WritDate    time.Time
}


func GetAllArtical() ([]Artical, error){
	orm := getLink()
	var articalModel []Artical
	err := orm.FindAll(&articalModel)

	return articalModel,err

}
