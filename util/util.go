package util

import (
	"math/rand"
	"time"
)

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
