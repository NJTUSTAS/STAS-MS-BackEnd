package common

import (
	"DemoProjectGO/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("What does this did?")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func GetToken(User model.User) (string, error) {
	//token过期时间：3小时后
	expirationTime := time.Now().Add(3 * time.Hour)
	//字段编写
	claims := &Claims{
		UserID: User.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "STAS-backend",
			Subject:   "user token",
		},
	}
	//生成token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	return token, err
}

// 原版返回了错误但是是空值，所以我决定不返回错误，让他报错吧。
func ParseToken(sToken string) (*jwt.Token, *Claims) {
	claims := &Claims{}
	// 无视错误，我也不知道有啥错误，所以不处理
	token, _ := jwt.ParseWithClaims(sToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims
}
