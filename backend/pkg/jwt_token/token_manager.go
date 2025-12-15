package jwttoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	SecretKey string        `envconfig:"JWT_SECRET_KEY"    required:"true"`
	ExpressIn time.Duration `envconfig:"JWT_EXPRESS_IN"    required:"true"`
	Issuer    string        `envconfig:"JWT_ISSUER"    required:"true"`
}

type JWTManager struct {
	SecretKey []byte
	ExpressIn time.Duration
	Issuer    string
}

func NewJWTManager(c Config) *JWTManager {
	return &JWTManager{
		SecretKey: []byte(c.SecretKey),
		ExpressIn: c.ExpressIn,
		Issuer:    c.Issuer,
	}
}

type Claims struct {
	UserID   int    `json:"user_id"`
	Usertag  string `json:"usertag"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (manager *JWTManager) CreateToken(userID int, usertag, username string) (string, error) {
	expirationTime := time.Now().Add(manager.ExpressIn)

	claims := Claims{
		UserID:   userID,
		Usertag:  usertag,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    manager.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(manager.SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func (manager *JWTManager) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return manager.SecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
