package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("nandarusfikri")

type JWTStruct struct {
	Authorized bool   `json:"authorize"`
	Client     string `json:"client"`
	Exp        string `json:"exp"`
}

func Auth(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if token != nil && err == nil {
	} else {
		result := gin.H{
			"message":    "not authorized",
			"error":      err.Error(),
			"api_status": 0,
		}
		c.JSON(401, result)
		c.Abort()
	}

}
