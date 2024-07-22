package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenType int

// token的状态.
const (
	ValidateToken TokenType = 0
	BadToken      TokenType = 1
	ExpiredToken  TokenType = 2
)

const AudienceUser = "user"
const AudienceAdmin = "admin"

type Claims struct {
	ID       int    `json:"id"`
	Role     string `json:"role"`
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

func GetJwtToken(secretKey, issuer string, seconds int64, claims Claims) (string, error) {
	expireTime := time.Now().Add(time.Duration(seconds) * time.Second)
	claims.ExpiresAt = jwt.NewNumericDate(expireTime)
	claims.Issuer = issuer

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := aToken.SignedString([]byte(secretKey))

	return token, err
}

func ParseToken(token string, secret string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		strings.TrimPrefix(token, "Bearer "), &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			return claims, nil
		}
	}
	return nil, err
}

func IsValidateOfToken(token string, secret string) (TokenType, *Claims) {
	claims, err := ParseToken(token, secret)
	if err != nil {
		return BadToken, nil
	}
	if !claims.VerifyExpiresAt(time.Now(), true) {
		return ExpiredToken, nil
	}
	if err := claims.Valid(); err != nil {
		return BadToken, nil
	}
	return ValidateToken, claims
}
