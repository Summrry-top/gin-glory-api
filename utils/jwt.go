package utils

import (
	"errors"
	"github.com/Summrry-top/gin-glory-api/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Account string
	IsAdmin bool
	jwt.StandardClaims
}

// 生成token
func CreateToken(account string, isAdmin bool) (string, error) {
	claims := MyClaims{
		Account: account,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(global.ServerConfig.Jwt.Expire).Unix(),
			Issuer:    global.ServerConfig.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 第一部分method，第二部分claims
	tokenStr, err := token.SignedString(global.JwtSecret)      // 第三部分秘钥
	return tokenStr, err
}

// 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return global.JwtSecret, nil
	})
	if err != nil {
		// 解析失败
		return nil, err
	}
	// 断言
	claims, ok := token.Claims.(*MyClaims)
	// 验证
	if !ok || !token.Valid {
		return nil, errors.New("JWT验证失败！")
	}
	return claims, nil
	// if token != nil {
	// 	// "正确的token字符串 可能有效，可能无效"
	// 	// 从token中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims 验证是否是我们定义的Claims和签名
	// 	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	// 	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
	// 		return claims, err
	// 	}
	// }
	// // "错误或无效的token字符串"
	// return nil, err
}
