package controller

import (
	"DemoProjectGO/database"
	"DemoProjectGO/model"
	"DemoProjectGO/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//业务逻辑
func Register(context *gin.Context) {
	//注册提供【用户名，邮箱，哈希密码三个参数】
	//用户名为空则随机生成用户名
	//邮箱不能重复
	db := database.GetDB()

	//从请求中获取数据。前端往后端请求的时候密码应该做一次哈希，因此这里直接用哈希后的密码。
	name := context.PostForm("name")
	email := context.PostForm("email")
	password := context.PostForm("password")
	log.Println("注册密码", password)
	passwordHashed, _ := util.Hash(password)
	//如果出错这里要返回http500内部错误并且返回，但是懒得写了

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
	if GetUserformEmail(db, email).ID != 0 {
		context.JSON(422, gin.H{
			"code": 422,
			"msg":  "exist email address!"})
		log.Println("邮箱已经注册过")
		//直接return，不进行数据库写入操作。
		return
	}
	//验证密码合法性应该在前端完成，不应该归后端管。
	//验证用户名
	if len(name) == 0 {
		//允许不取名，系统生成16位随机16进制字符。
		name = util.RandomHexString(16)
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
	//默认行为：创建数据库
	//注意这里要传引用
	db.Create(&model.User{Name: name, Email: email, Hashword: passwordHashed})
}

func Login(context *gin.Context) {
	//提供邮箱和密码；
	//邮箱应当存在，否则报错
	//密码应当匹配。否则报错
	//返回token
	db := database.GetDB()

	//从请求中获取数据。前端往后端请求的时候密码应该做一次哈希，因此这里直接用哈希后的密码。
	email := context.PostForm("email")
	password := context.PostForm("password")
	log.Println("登录密码", password)
	passwordHashed, _ := util.Hash(password)
	log.Println("登录hash：", passwordHashed)
	//passwordHashed := Hash(password, context)

	//合法性验证由前端完成，进行用户存在性验证
	user := GetUserformEmail(db, email)
	if user.ID == 0 {
		log.Println("用户不存在")
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//密码匹配验证
	if !util.PasswordMatchQ(password, user.Hashword) {
		log.Println("密码不匹配")
		context.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户名与密码不匹配"})
		return
	}

	//默认正常行为：发放token
	token := util.RandomHexString(16)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})

}

func GetUserformEmail(db *gorm.DB, email string) model.User {
	//不存在为0
	var user model.User
	//查找数据库并且把找到的第一个结果传给user
	db.Where("email = ?", email).First(&user)
	//user.ID是在默认值里的，如果找不到那么ID就是0
	return user
}
