package model

import "github.com/jinzhu/gorm"

//数据库读写用
type User struct {
	//第一项是默认的，编号，创建时间，编辑时间、删除时间，不过我还不知道删除时间有啥用
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(110);not null;unique"`
	Hashword string `gorm:"size:255;not null"`
}
