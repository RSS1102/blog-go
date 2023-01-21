package Jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Name                 string `json:"name"`
	jwt.RegisteredClaims        //jwt自带载荷
}

var mySigningKey = []byte("AllYourBase")

func GenerateClaims(name string) Claims {
	return Claims{
		name,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //生效时间
			Issuer:    "Jimmy",                                            //签发者
			Subject:   "bolg",                                             //签发主题
			ID:        "1",                                                //
			Audience:  []string{name},                                     //接收者
		},
	}
}

func GenerateJwt(name string) (string, error) {
	// Create the claims
	claims := GenerateClaims(name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func AnalysisJwt(tokenString string) (*Claims, error) {
	name := new(Claims)
	claims, err := jwt.ParseWithClaims(tokenString, name, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token error:%v", err)
	}
	return name, err
}
