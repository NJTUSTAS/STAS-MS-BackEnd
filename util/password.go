package util

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// GetPassword 获取密码，密码和盐用:分隔
func GetPassword(plainText string) string {
	salt := RandomHexString(3)
	return fmt.Sprintf("%x:%s", hash(plainText, salt), salt)
}

func hash(plainText, salt string) []byte {
	return pbkdf2.Key([]byte(plainText), []byte(salt), 4096, sha256.Size, sha256.New)
}

// CheckPassword 检测密码是否正确
func CheckPassword(plainText, password string) bool {
	salt := strings.Split(password, ":")[1]
	return password == fmt.Sprintf("%x:%s", hash(plainText, salt), salt)
}
