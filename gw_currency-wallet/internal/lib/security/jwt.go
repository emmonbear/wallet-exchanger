package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emmonbear/wallet-exchanger/internal/model"
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
