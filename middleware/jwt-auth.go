package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	service "github.com/fabiosebastiano/jwt-golang"
	"github.com/gin-gonic/gin"
)

const (
	ISSUER     string        = "fabiosebastiano.com"
	SECRET     string        = "secret"
	EXPIRATION time.Duration = 1
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService(ISSUER, SECRET, EXPIRATION).ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
