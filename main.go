package main

//gin，gorm和mysql都要用gin安装
import (
	"fmt"
	"github.com/gin-gonic/gin"
	//这里要加一行来完成对mysql的驱动初始化
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"time"
)

//数据库读写用
type User struct {
	//第一项是默认的，编号，创建时间，编辑时间、删除时间，不过我还不知道删除时间有啥用
	gorm.Model
	Name           string `gorm:"type:varchar(20);not null"`
	Email          string `gorm:"type:varchar(110);not null;unique"`
	PasswordHashed string `gorm:"size:255;not null"`
}

func main() {
	fmt.Println("start working...")

	//首先连接数据库并且在离开函数之前关闭连接（使用defer）
	//defer 会在return之后调用，类似with open() as xxx这样
	//数据库信息硬编码在InitDatabase()里，可能也有不好的地方但是也能用
	db := InitDatabase()
	defer db.Close()
	fmt.Println("database connected")

	//以下开始写接口
	router := gin.Default()
	//第一个参数：网站相对地址，第二个参数是个很大的函数，也即业务逻辑
	router.POST(
		"/api/auth/register",
		//以下是第二个参数的范畴。函数和对应的请求分开没有显然的好处并且造成很多麻烦，所以写在一起
		func(context *gin.Context) {
			//从请求中获取数据。前端往后端请求的时候密码应该做一次哈希，因此这里直接用哈希后的密码。
			name := context.PostForm("name")
			email := context.PostForm("email")
			passwordHashed := context.PostForm("password")

			//以下开始验证
			//邮箱合法性验证
			if len(email) ==0 {
				//这里假设只要求非空
				context.JSON(422, gin.H{
					"code": 422,
					"msg":  "illegal email address!"})
				log.Printf("非法邮箱：%s，注册失败",email)
				//直接return，不进行数据库写入操作。
				return
			}
			//电话重复性验证
			if EmailExistQ(db, email) {
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
				name = RandomHexName(16)
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
			newUser :=User{Name: name, Email: email, PasswordHashed:passwordHashed}
			log.Println("开始写入数据库")
			//注意这里要传引用
			db.Create(&newUser)
			log.Println("结束写入数据库")


		},
		//第二行传递的函数到此为止
	)
	//post 端口到此结束

	//上面是业务逻辑，现在开始运行，默认端口是8080
	_ = router.Run(":8080")
}

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
	db.AutoMigrate(&User{})
	return db
}

func EmailExistQ(db *gorm.DB, email string) bool {
	var user User
	//查找数据库并且把找到的第一个结果传给user
	db.Where("email = ?", email).First(&user)
	//user.ID是在默认值里的，如果找不到那么ID就是0
	return user.ID != 0
}

//随机16进制字符实现，对着视频里的凑合弄了一下，写的有点蠢
func RandomHexName(n int) string {
	//初始化随机数池，理论上来说这个不应该每次调用生成一次，最好用静态的或者全局？
	var chars = []byte("0123456789ABCDFEF")
	//申请内存
	retB := make([]byte, n)
	//初始化种子
	rand.Seed(time.Now().Unix())
	//逐个生成随机数，好蠢啊
	for i := range retB {
		retB[i] = chars[rand.Intn(len(chars))]
	}
	return string(retB)
}
