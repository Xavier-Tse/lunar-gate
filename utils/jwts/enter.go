package jwts

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/Xavier-Tse/lunar-gate/global"
	"time"
)

type ClaimsUserInfo struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	RoleList []uint `json:"role"`
}

type Claims struct {
	ClaimsUserInfo
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(info ClaimsUserInfo) (string, error) {
	j := global.Config.Jwt
	cla := Claims{
		ClaimsUserInfo: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(j.Expires) * time.Hour).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	return token.SignedString([]byte(j.Secret))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	j := global.Config.Jwt
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		if claims.Issuer != j.Issuer {
			return nil, errors.New("invalid issuer")
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
