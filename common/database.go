package common

import (
	"DemoProjectGO/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	//这里要加一行来完成对mysql的驱动初始化
	_ "github.com/go-sql-driver/mysql"
)

//var db *gorm.DB
var db *gorm.DB = InitDatabase()

func InitDatabase() *gorm.DB {
	//数据库类型，地址，端口，地址，数据库名称，密码
	driverName := "mysql"
	host := "106.13.162.70"
	port := "3306"
	database := "users"
	username := "rdsroot"
	password := "qwer1234"

	//编码。教程用的是utf8，但是sql的utf8是假的utf8,utf8mb4才是真utf8
	charset := "utf8mb4"

	//填充数据
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	//_忽略错误，正常情况下第二项应该是nil，否则提示报错信息
	db, err := gorm.Open(driverName, args)
	//输出错误信息
	if err != nil {
		panic(err)
	}
	//自动生成数据表(不存在即创建)
	db.AutoMigrate(&model.User{})
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Printf("你妈的为什么 给我一个空指针")
	}
	return db
}
