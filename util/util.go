package util

import (
	"crypto/rand"
	"fmt"
)

// RandomHexString 生成n位的随机字符串（16进制）
func RandomHexString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
