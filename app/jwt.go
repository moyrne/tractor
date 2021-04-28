package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/moyrne/tractor/configs"
	"time"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(configs.ServerSetting.JWTSecret)
}

func GenerateToken(userId uint) (string, error) {
	claims := Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(configs.ServerSetting.JWTExpire)).Unix(),
			Issuer:    configs.ServerSetting.JWTIssuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims == nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if ok && tokenClaims.Valid {
		return claims, nil
	}

	return claims, err
}
