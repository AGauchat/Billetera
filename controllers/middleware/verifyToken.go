package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyToken(c *gin.Context) {
	coockieToken, err := c.Cookie("token")
	headerToken := c.Request.Header.Get("token")

	if err != nil {
		if err == http.ErrNoCookie {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if coockieToken != headerToken {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("bille"), err
	})

	if token.Valid {
		Refresh(c)
		c.Next()
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.AbortWithStatus(http.StatusBadRequest)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			c.AbortWithStatus(http.StatusRequestTimeout)
		} else {
			c.AbortWithStatus(400)
		}
	} else {
		c.AbortWithStatus(400)
	}
}
