package services

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type TockenService struct {
	ctx       context.Context
	secretKey []byte
}

func NewTockenService(secretKey []byte) *TockenService {
	return &TockenService{
		context.Background(),
		secretKey,
	}
}

func (s *TockenService) SetContext(context context.Context) *TockenService {
	s.ctx = context
	return s
}

func (s *TockenService) VerifyToken(tocken string) (*jwt.Token, jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tocken, claims, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	return token, claims, err
}

func (s *TockenService) VerifyTokenWithSecret(tocken string, secret []byte) (*jwt.Token, jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tocken, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	return token, claims, nil
}

func (s *TockenService) SignToken(data map[string]interface{}) (*string, error) {
	var claims jwt.MapClaims = data

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (s *TockenService) SignTokenWithSecret(data map[string]interface{}, secret []byte) (*string, error) {
	var claims jwt.MapClaims = data

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
