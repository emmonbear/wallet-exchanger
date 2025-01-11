package security

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/model"
)

const tokenTTL = 12 * time.Hour
const signingKey = "aqwec@4qvhl;#2;lg"

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateJWTToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseJWTToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
