package util

import (
	"DemoProjectGO/common"
	"DemoProjectGO/model"
	"crypto/rand"
	"fmt"
)

// RandomHexString 生成n位的随机字符串（16进制）
func RandomHexString(n int) string {
	randBytes := make([]byte, n/2)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// GetUserFormEmail works as its name
func GetUserFormEmail(email string) model.User {
	//不存在为0
	var user model.User
	//查找数据库并且把找到的第一个结果传给user
	common.GetDB().Where("email = ?", email).First(&user)
	//user.ID是在默认值里的，如果找不到那么ID就是0
	return user
}

// GetUserFormID works as its name
func GetUserFormID(userID uint) model.User {
	//不存在为0
	var user model.User
	//查找数据库并且把找到的第一个结果传给user
	common.GetDB().First(&user, userID)
	//user.ID是在默认值里的，如果找不到那么ID就是0
	return user
}
