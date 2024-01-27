package access

import (
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	AccessUuid string `json:"accessUuid"`
	UserUuid   string `json:"userUuid"`
	UserRole   string `json:"userRole"`
	Exp        int    `json:"exp"`
	jwt.StandardClaims
}

type TokenService struct {
	accessSecret string
}

func NewTokenService(accessSecret string) *TokenService {
	return &TokenService{
		accessSecret: accessSecret,
	}
}

func (r *TokenService) GetAccessSecret() string {
	return r.accessSecret
}

func (r *TokenService) DecodeAccessToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.accessSecret), nil
	})
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
