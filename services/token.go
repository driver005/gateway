package services

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("TODO: Change me")

type TockenService struct {
	ctx context.Context
}

func NewTockenService(ctx context.Context) *TockenService {
	return &TockenService{
		ctx,
	}
}

func (s *TockenService) VerifyToken(tocken string) (*jwt.Token, jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tocken, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	return token, claims, err
}

func (s *TockenService) VerifyTokenWithSecret(tocken string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tocken, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func (s *TockenService) SignToken(data map[string]interface{}) (*string, error) {
	var claims jwt.MapClaims = data

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
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
