package identity

import (
	"errors"
	"hundis/config"
	"hundis/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID uint `json:"userId"`
}

func CreateJWT(user *model.User) (string, error) {
	config := config.Config()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (uint, error) {
	config := config.Config()
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		if claims.UserID > 0 {
			return claims.UserID, nil
		}
		return 0, errors.New("invalid userId in token")
	}

	if !token.Valid {
		return 0, jwt.ErrTokenExpired
	}

	return 0, jwt.ErrTokenUnverifiable
}
