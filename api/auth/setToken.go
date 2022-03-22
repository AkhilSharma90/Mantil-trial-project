package auth

import (
	"time"
	"context"
	"github.com/dgrijalva/jwt-go"
)


func (t *Token) SetToken(ctx context.Context) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	
	claims := Claims{
		"akhil",
		"akhilsails@gmail.com",
		"788949993",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "akhil sharma",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))


	return signedToken
}