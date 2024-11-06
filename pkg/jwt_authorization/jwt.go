package jwt_authorization

import (
	"github.com/FcorpionItsMe/ftodo/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TODO: Потом реализовать ключ в Env
var secret = []byte("super secret! Dont say to anybody!")

type JWTAuth struct {
}

func NewJWTAuth() *JWTAuth {
	return &JWTAuth{}
}

func (j JWTAuth) CreateToken(userInfo domain.SignInUserInput) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userInfo.Login,
		"exp":      time.Now().Add(time.Minute * 20).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (j JWTAuth) ParseToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, nil
	}
	return true, nil
}
