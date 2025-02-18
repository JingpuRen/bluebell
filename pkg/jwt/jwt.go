package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TokenExpireDuration = time.Hour * 2

var salty = []byte("JingpuRen")

type MyClaims struct {
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌的标准声明
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		userID,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "bluebell", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象，第一个参数是加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的盐值并获得完整的编码后的字符串token
	return token.SignedString(salty)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return salty, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
