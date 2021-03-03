package model

import "github.com/jinzhu/gorm"

type Fresh struct { //数据格式
	gorm.Model
	Name         string `gorm:"type:varchar(40);not null"` //姓名
	Major        string `gorm:"type:varchar(20)"`          //专业
	StudentId    string `gorm:"type:varchar(15);not null"` //学号
	Phone        string `gorm:"type:varchar(11);not null"` //手机
	Grade        string `gorm:"type:varchar(4)"`           //年级
	Gender       string `gorm:"type:varchar(4)"`           //性别
	FirstChoice  string `gorm:"type:varchar(10);not null"` //第一志愿
	SecondChoice string `gorm:"type:varchar(10)"`          //第二志愿
	Introduction string `gorm:"type:varchar(500)"`         //自我介绍
	Hope         string `gorm:"type:varchar(500)"`         //期望
	Hobbies      string `gorm:"type:varchar(100)"`         //兴趣
}
