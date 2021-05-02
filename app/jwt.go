package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/moyrne/tractor/configs"
	"github.com/moyrne/tractor/errcode"
	"net/http"
	"time"
)

type Claims struct {
	CustomValue CustomValue `json:"value"`
	jwt.StandardClaims
}

type CustomValue map[string]interface{}

func (c *CustomValue) SetValue(key string, value interface{}) {
	(*c)[key] = value
}

var ErrJWTAssert = errcode.NewError(http.StatusInternalServerError, "jwt assert error")

func (c *CustomValue) GetInt(key string) (int, error) {
	v, ok := (*c)[key].(int)
	if !ok {
		return 0, errcode.NewErrInfo(ErrJWTAssert, "auth", "GetInt error")
	}
	return v, nil
}

func (c *CustomValue) GetString(key string) (string, error) {
	v, ok := (*c)[key].(string)
	if !ok {
		return "", errcode.NewErrInfo(ErrJWTAssert, "auth", "GetString error")
	}
	return v, nil
}

func (c *CustomValue) GetBytes(key string) ([]byte, error) {
	v, ok := (*c)[key].([]byte)
	if !ok {
		return nil, errcode.NewErrInfo(ErrJWTAssert, "auth", "GetBytes error")
	}
	return v, nil
}

func (c *CustomValue) GetIntP(key string) int {
	v, err := c.GetInt(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CustomValue) GetStringP(key string) string {
	v, err := c.GetString(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CustomValue) GetBytesP(key string) []byte {
	v, err := c.GetBytes(key)
	if err != nil {
		panic(err)
	}
	return v
}

func GenerateToken(customValue CustomValue) (string, error) {
	claims := Claims{
		CustomValue: customValue,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(configs.JWTSetting.Expire)).Unix(),
			Issuer:    configs.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokenClaims.SignedString(configs.JWTSetting.Secret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return configs.JWTSetting.Secret, nil
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
