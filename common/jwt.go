package common

import (
	"DemoProjectGO/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//我并不知道这个c字符串起到什么作用。
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
