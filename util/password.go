package util

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// GetPasswordWithSalt 获取密码，密码和盐用:分隔
func GetPasswordWithSalt(plainPassword string) string {
	salt := RandomHexString(3)
	return fmt.Sprintf("%x:%s", hash(plainPassword, salt), salt)
}

func hash(plainPassword, salt string) []byte {
	return pbkdf2.Key([]byte(plainPassword), []byte(salt), 4096, sha256.Size, sha256.New)
}

// CheckPassword 检测密码是否正确
func CheckPassword(plainPassword, passwordWithSalt string) bool {
	salt := strings.Split(passwordWithSalt, ":")[1]
	return passwordWithSalt == fmt.Sprintf("%x:%s", hash(plainPassword, salt), salt)
}
