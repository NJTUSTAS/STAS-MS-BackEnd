package controller

import (
	"DemoProjectGO/common"
	"DemoProjectGO/model"
	"DemoProjectGO/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//业务逻辑

//Register work as its name
func Register(context *gin.Context) {
	//注册提供【用户名，邮箱，哈希密码三个参数】
	//用户名为空则随机生成用户名
	//邮箱不能重复
	db := common.GetDB()

	//从请求中获取数据。前端往后端请求的时候密码应该做一次哈希，因此这里直接用哈希后的密码。
	name := context.PostForm("Name")
	email := context.PostForm("Email")
	password := context.PostForm("Password")
	passwordHashed, _ := util.Hash(password)
	//如果出错这里要返回http500内部错误并且返回，但是懒得写了

	//以下开始验证
	//邮箱合法性验证
	if len(email) == 0 {
		//这里假设只要求非空
		fmt.Println("print email:", email)
		context.JSON(422, gin.H{
			"code": 422,
			"msg":  "illegal email address!"})
		//log.Printf("非法邮箱：%s，注册失败", email)
		//直接return，不进行数据库写入操作。
		return
	}
	//邮箱重复性验证
	if util.GetUserFormEmail(email).ID != 0 {
		context.JSON(422, gin.H{
			"code": 422,
			"msg":  "exist email address!"})
		//log.Println("邮箱已经注册过")
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
		//log.Printf("无用户名注册成功，生成用户名：%s", name)
	} else {
		//有用户名，成功注册
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "register successful."})
		//log.Println("注册成功")
	}
	//默认行为：创建数据库
	//注意这里要传引用
	user := model.User{Name: name, Email: email, Hashword: passwordHashed}
	db.Create(&user)
}

//Login work as its name
func Login(context *gin.Context) {
	//提供邮箱和密码；
	//邮箱应当存在，否则报错
	//密码应当匹配。否则报错
	//返回token

	//从请求中获取数据。前端往后端请求的时候密码应该做一次哈希，因此这里直接用哈希后的密码。
	email := context.PostForm("Email")
	password := context.PostForm("Password")

	//合法性验证由前端完成，进行用户存在性验证
	user := util.GetUserFormEmail(email)
	if user.ID == 0 {
		//log.Println("用户不存在")
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//密码匹配验证
	if !util.PasswordMatchQ(password, user.Hashword) {
		//log.Println("密码不匹配")
		context.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户名与密码不匹配"})
		return
	}

	//默认正常行为：发放token
	token, err := common.GetToken(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token err:%v", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

//EnrollReceive 招新收集表
func EnrollReceive(context *gin.Context) {
	//目前没弄清楚context.Request.PostFormValue和context.PostForm之间有啥区别
	//读入数据
	cRP := context.Request.PostFormValue
	name := cRP("name")
	studentID := cRP("studentId")
	major := cRP("major")
	phone := cRP("phone")
	grade := cRP("grade")
	gender := cRP("gender")
	firstChoice := cRP("firstChoice")
	secondChoice := cRP("secondChoice")
	introduction := cRP("introduction")
	hope := cRP("hope")
	hobbies := cRP("hobbies")

	newFreshman := model.Fresh{
		Name:         name,
		StudentId:    studentID,
		Major:        major,
		Phone:        phone,
		Grade:        grade,
		Gender:       gender,
		FirstChoice:  firstChoice,
		SecondChoice: secondChoice,
		Introduction: introduction,
		Hope:         hope,
		Hobbies:      hobbies,
	}

	db := common.GetDB()
	db.Create(&newFreshman)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "报名成功",
	})
}

// Info to get user info
func Info(context *gin.Context) {
	user, exist := context.Get("user")
	fmt.Println("user,exist=", user, exist)
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "not login yet."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": util.ToUserOutput(user.(model.User))}})
}
