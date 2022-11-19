package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-campaign-app/helper"
)

type JWTAuth interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTAuthImpl struct {
}

var key = []byte("S3creT K3y")

func (auth *JWTAuthImpl) ValidateToken(token string) (*jwt.Token, error) {
	parse, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("Invalid token")
			}
			return key, nil
		})

	if err != nil {
		return nil, err
	}

	return parse, nil
}

func (auth *JWTAuthImpl) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedString, err := token.SignedString(key)
	helper.PanicIfError(err)

	return signedString, nil
}

func NewJWTAuthImpl() JWTAuth {
	return &JWTAuthImpl{}
}
