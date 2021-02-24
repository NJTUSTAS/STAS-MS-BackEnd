package controller

import (
	"DemoProjectGO/common"
	"DemoProjectGO/model"
	"DemoProjectGO/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

//业务逻辑
func Register(context *gin.Context) {
	//注册提供【用户名，邮箱，哈希密码三个参数】
	//用户名为空则随机生成用户名
	//邮箱不能重复
	//哈希密码为密码的哈希值
	db := common.GetDB()

	//从请求中获取数据。前端往后端请求的时候密码应该做一次哈希，因此这里直接用哈希后的密码。
	name := context.PostForm("name")
	email := context.PostForm("email")
	passwordHashed := context.PostForm("password")

	//以下开始验证
	//邮箱合法性验证
	if len(email) == 0 {
		//这里假设只要求非空
		context.JSON(422, gin.H{
			"code": 422,
			"msg":  "illegal email address!"})
		log.Printf("非法邮箱：%s，注册失败", email)
		//直接return，不进行数据库写入操作。
		return
	}
	//电话重复性验证
	if GetIDformEmail(db, email) != 0 {
		context.JSON(422, gin.H{
			"code": 422,
			"msg":  "exist email address!"})
		log.Println("邮箱已经注册过")
		//直接return，不进行数据库写入操作。
		return
	}
	//验证密码应该在前端完成，不应该归后端管。
	//验证用户名
	if len(name) == 0 {
		//允许不取名，系统生成16位随机16进制字符。
		name = util.RandomHexName(16)
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "no user name,auto generated.",
			"name": name})
		log.Printf("无用户名注册成功，生成用户名：%s", name)
	} else {
		//有用户名，成功注册
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "register successful."})
		log.Println("注册成功")
	}

	//通过验证，可以开始写入了。先生成数据结构，在表中创建对应的行。
	newUser := model.User{Name: name, Email: email, PasswordHashed: passwordHashed}
	log.Println("开始写入数据库")
	//注意这里要传引用
	db.Create(&newUser)
	log.Println("结束写入数据库")
}


func GetIDformEmail(db *gorm.DB, email string) uint {
	//不存在为0
	var user model.User
	//查找数据库并且把找到的第一个结果传给user
	db.Where("email = ?", email).First(&user)
	//user.ID是在默认值里的，如果找不到那么ID就是0
	return user.ID
}
