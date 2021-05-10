package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Refresh(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	headerToken := c.Request.Header.Get("token")
	if cookie != headerToken {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tknStr := cookie
	claims := &Claim{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bille"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("bille"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", tokenString, 360, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
