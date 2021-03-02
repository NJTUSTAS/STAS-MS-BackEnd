import (
	"fmt"
	"github.com/gin-gonic/gin"
	//首先在这里添加初始化数据库的包，不然无法运行
)

type Freshman struct { //数据格式
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null"` //姓名
	Major        string `gorm:"type:varchar(15)"`          //专业
	Id           string `gorm:"type:varchar(15);not null"` //学号
	Phone        string `gorm:"type:varchar(11);not null"` //电话
	Grade        string `gorm:"type:varchar(5)"`           //年级
	Gendar       string `gorm:"type:varchar(2)"`           //性别
	FirstWish    string `gorm:"type:varchar(10);not null"` //第一志愿
	SecondWish   string `gorm:"type:varchar(10)"`          //第二志愿
	Introduction string `gorm:"type:varchar(200)"`         //自我介绍
	Hope         string `gorm:"type:varchar(200)"`         //期望
	Hobbies      string `gorm:"type:varchar(200)"`         //兴趣
}

func Recruit_new bool {//创建成功返回true
	db := InitDb()
	db.AutoMigrate(&Freshman{}) //没有则自动创建表，如果报错panic: Error 1146: Table '数据库.freshmen' doesn't exist请手动创建一个表名为freshmen的空表
	defer db.Close()
	r := gin.Default()
	r.POST("/receive"/*网页的路径，测试时改为正式网址*/, func(c *gin.Context) { //接受参数，post传递，以表单形式form-data或x-www-form-urlencoded
		name := c.Request.PostFormValue("name")
		id := c.Request.PostFormValue("id")
		major := c.Request.PostFormValue("major")
		phone := c.Request.PostFormValue("phone")
		grade := c.Request.PostFormValue("grade")
		gendar := c.Request.PostFormValue("gendar")
		firstWish := c.Request.PostFormValue("firstWish")
		secondWish := c.Request.PostFormValue("secondWish")
		introduction := c.Request.PostFormValue("introduction")
		hope := c.Request.PostFormValue("hope")
		hobbies := c.Request.PostFormValue("hobbies")
		
		newFreshman := Freshman{
			Name:         name,
			Id:           id,
			Major:        major,
			Phone:        phone,
			Grade:        grade,
			Gendar:       gendar,
			FirstWish:    firstWish,
			SecondWish:   secondWish,
			Introduction: introduction,
			Hope:         hope,
			Hobbies:      hobbies,
		}
		db.Create(&newFreshman)
		if db.Error != nil {
			return false
		} else {
			return true
		}
	})
	panic(r.Run())


}