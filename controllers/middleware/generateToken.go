package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(idUsuario int64) string {
	//expirationTime := time.Now().Add(5 * time.Minute)
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := Claim{
		idUsuario,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tok.SignedString([]byte("bille"))

	if err != nil {
		return "No se pudo generar el middleware."
	}

	return token
}
