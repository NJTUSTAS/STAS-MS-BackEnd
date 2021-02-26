package util

import (
	"golang.org/x/crypto/bcrypt"
)

//// GetPasswordWithSalt 获取密码，密码和盐用:分隔
//func GetPasswordWithSalt(plainPassword string) string {
//	salt := RandomHexString(3)
//	return fmt.Sprintf("%x:%s", hash(plainPassword, salt), salt)
//}

func Hash(plainPassword string) (string, error) {
	cRet, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	//c表示是[]byte 类型，即c风格字符串
	return string(cRet), err
}

// PasswordMatchQ 检测密码是否匹配
func PasswordMatchQ(plainPassword, passwordWithSalt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordWithSalt), []byte(plainPassword))
	return err == nil
}
