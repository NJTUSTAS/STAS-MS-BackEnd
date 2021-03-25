package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Name string`gorm:"type:varchar(10);not null"`
	Data string`gorm:"type:varchar(200);not null"`
}
