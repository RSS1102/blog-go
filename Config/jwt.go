package Jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        //jwt自带载荷
}

var mySigningKey = []byte("cipher")

// GenerateClaims 生成Claims
func customClaims(username string) Claims {
	return Claims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //生效时间
			Issuer:    "Jimmy",                                            //签发者
			Subject:   "bolg",                                             //签发主题
			ID:        "1",                                                //
			Audience:  []string{username},                                 //接收者
		},
	}
}

// GenerateJwt 生成Token
func GenerateJwt(username string) (string, error) {
	// Create the claims
	claims := customClaims(username)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalysisJwt 解析Token
func AnalysisJwt(tokenString string) (*Claims, error) {
	customClaims := new(Claims)
	// 解析token
	claims, err := jwt.ParseWithClaims(tokenString, customClaims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token error:%v", err)
	}
	return customClaims, err
}
