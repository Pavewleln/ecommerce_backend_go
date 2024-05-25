package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString != "" {
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
			})

			if err != nil || !token.Valid {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
				return
			}

			userID := claims["_id"].(string)
			ctx := context.WithValue(c.Request.Context(), "user_id", userID)
			c.Request = c.Request.WithContext(ctx)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			return
		}

		c.Next()
	}
}
