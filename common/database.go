package common

import (
	"DemoProjectGO/model"
	"fmt"

	"github.com/jinzhu/gorm"

	//这里要加一行来完成对mysql的驱动初始化
	_ "github.com/go-sql-driver/mysql"
)

//var db *gorm.DB
var db *gorm.DB

// InitDatabase
func InitDatabase() *gorm.DB {
	//数据库类型，地址，端口，地址，数据库名称，密码
	driverName := "mysql"
	port := "3306"

	//丁昊的rds
	//host := "106.13.162.70"
	//database := "users"
	//username := "rdsroot"
	//password := "qwer1234"

	//杨凯的服务器
	host := "192.144.128.226"
	database := "users"
	username := "root"
	password := "123456"

	//编码。教程用的是utf8，但是sql的utf8是假的utf8,utf8mb4才是真utf8
	charset := "utf8mb4"

	//填充数据
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic(err)
	}

	//自动生成数据表(不存在即创建)
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Fresh{})
	db.AutoMigrate(&model.Note{})

	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		//没有就自动生成
		db = InitDatabase()
	}
	return db
}
