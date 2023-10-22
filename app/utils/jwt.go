package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
type JwtTokenType string

const (
	AccessToken  JwtTokenType = "access_token_exp"
	RefreshToken JwtTokenType = "refresh_token_exp"
)

func GenerateJWTToken(userId uint, tokenType JwtTokenType) (string, error) {

	var expirationTime time.Time

	if tokenType == AccessToken {
		expirationTime = time.Now().Add(15 * time.Minute)
	} else if tokenType == RefreshToken {
		expirationTime = time.Now().Add(168 * time.Hour)
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	claims := &Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Println("Happened error when generate token.")
		return "", err
	}

	return tokenString, nil
}

func ParseJWTToken(tokenString string) (*Claims, error) {

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Println("Happened error when parse token. Error:", err.Error())
		return nil, err
	}

	if !token.Valid {
		log.Println("Happened error when validate token.")
		return nil, err
	}

	return claims, nil
}
