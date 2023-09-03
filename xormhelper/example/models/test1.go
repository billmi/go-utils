package models

type Test1 struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Title      string `json:"title" xorm:"VARCHAR(50)"`
	Status     int    `json:"status" xorm:"TINYINT(4)"`
	CreateTime int    `json:"create_time" xorm:"INT(11)"`
	UpdateTime int    `json:"update_time" xorm:"INT(11)"`
}
