package main

//gin，gorm和mysql都要用gin安装
import (
	"DemoProjectGO/common"
	"DemoProjectGO/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start working...")

	//首先连接数据库并且在离开函数之前关闭连接（使用defer）
	//defer 会在return之后调用，类似with open() as xxx这样
	//数据库信息硬编码在InitDatabase()里，可能也有不好的地方但是也能用
	//db := common.InitDatabase()
	db := common.GetDB()
	defer db.Close()

	fmt.Println("database connected")

	//以下开始写接口
	router := gin.Default()
	//router = CollectRoute(router)
	router.POST("/api/auth/register", controller.Register)
	//post 端口到此结束

	//上面是业务逻辑，现在开始运行，默认端口是8080
	_ = router.Run(":8080")
}
